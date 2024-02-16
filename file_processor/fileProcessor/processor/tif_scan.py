import os
from PIL import Image, ImageSequence
from azure.storage.blob import BlobServiceClient
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobClient

from utils.utils import download_from_azure
from database import FileDatabase


class TifScan:
    def __init__(self, file_id):
        self.database = FileDatabase()
        self.halted = False
        self.file_id = file_id

        file = self.database.get_file_by_id(file_id)
        if file.status == "processed":
            self.halted = True
            return
        self.txt_file_name = file.name.replace(".tif", ".txt")

        txt_file = self.database.get_file_by_name_dataset(self.txt_file_name, file.dataset_id)
        if txt_file is None:
            self.database.update_status(self.file_id, "awaitingtxt")
            self.halted = True
        else:
            self.txt_file_id = txt_file.id

    def download_files(self):
        if self.halted:
            return False
        download_from_azure(self.file_id)
        download_from_azure(self.txt_file_id)

        return True

    def process(self):
        if self.halted:
            return False

        with open(".temp/"+self.txt_file_id, 'r') as file:
            for raw_line in file:
                line = raw_line.strip()
                if "=" not in line:
                    continue
                if "   " in line:
                    primary_attribute, child_attributes = line.split("=", 1)
                    attributes = child_attributes.split("   ")
                    for attribute in attributes:
                        key, value = attribute.split("=")
                        if value != "":
                            self.database.add_attribute(self.file_id, primary_attribute + "-" + key, value)
                    continue
                key, value = line.split("=")
                if value != "" or len(value) > 0:
                    self.database.add_attribute(self.file_id, key, value)

        self.get_preview()

        self.database.update_status(self.txt_file_id, "support_processed")
        return True

    def get_preview(self):
        if self.halted:
            return False
        image = Image.open(".temp/"+self.file_id)

        for i, page in enumerate(ImageSequence.Iterator(image)):
            try:
                page.save(".temp/"+self.file_id.split(".")[0] + ".png")
            except:
                print(page)
            break

        url = "https://synopticprojectstorage.blob.core.windows.net/"
        token_credential = DefaultAzureCredential()
        blob_service_client = BlobServiceClient(account_url=url, credential=token_credential)
        container_client = blob_service_client.get_container_client(container="fyp-previews")
        with open(file=".temp/"+self.file_id.split(".")[0] + ".png", mode="rb") as data:
            container_client.upload_blob(name=self.file_id.split(".")[0] + ".png", data=data, overwrite=True)

        os.remove(".temp/"+self.file_id.split(".")[0] + ".png")

    def finish_process(self):
        if self.halted:
            return False
        os.remove(".temp/"+self.file_id)
        os.remove(".temp/"+self.txt_file_id)
        return True