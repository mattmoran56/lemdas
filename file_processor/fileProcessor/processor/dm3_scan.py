import os
from PIL import Image, ImageSequence
from azure.storage.blob import BlobServiceClient
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobClient

from database import FileDatabase
from processor.scan import Scan
import hyperspy.api as hs

class DM3Scan(Scan):
    def __init__(self, file_id):
        self.database = FileDatabase()
        self.halted = False
        self.file_id = file_id

        file = self.database.get_file_by_id(file_id)
        if file.status == "processed":
            self.halted = True
            return

    def process(self):
        if self.halted:
            return False

        s = hs.load(".temp/"+self.file_id)
        metadata = s.metadata.as_dictionary()

        for (key, value) in metadata.items():
            if isinstance(value, dict):
                self.extract_metadata(value, key)
            else:
                self.database.add_attribute(self.file_id, key, value)

        return True

    def extract_metadata(self, metadata, prefix):
        for (key, value) in metadata.items():
            if isinstance(value, dict):
                self.extract_metadata(value, prefix + "-" + key)
            else:
                self.database.add_attribute(self.file_id, prefix + "-" + key, value)

    def get_preview(self):
        image = Image.open(".temp/" + self.file_id)

        for i, page in enumerate(ImageSequence.Iterator(image)):
            try:
                page.save(".temp/" + self.file_id.split(".")[0] + ".png")
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
