import os
from PIL import Image, ImageSequence
from azure.storage.blob import BlobServiceClient
from azure.identity import DefaultAzureCredential

from database import FileDatabase
from processor.scan import Scan
import hyperspy.api as hs

from PIL import Image
import numpy as np
import ncempy.io as nio

class DM3Scan(Scan):
    def __init__(self, file_id):
        self.database = FileDatabase()
        self.halted = False
        self.file_id = file_id

        file = self.database.get_file_by_id(file_id)
        if file.status == "processed":
            self.halted = True
            return

    def tuple_to_string(self, tuple):
        strVar = ""

        for i in tuple:
            strVar += str(i)
            strVar += " "

        return strVar

    def process(self):
        if self.halted:
            return False

        s = hs.load(".temp/"+self.file_id)
        metadata = s.original_metadata.as_dictionary()

        print(s.original_metadata)

        print(metadata)
        
        group_id = self.database.add_attribute_group(self.file_id, "rootgroup")

        for (key, value) in metadata.items():
            if isinstance(value, dict):
                self.extract_metadata(value, key, group_id)
            elif isinstance(value, tuple):
                self.database.add_attribute(self.file_id, key, self.tuple_to_string(value), group_id)
            else:
                self.database.add_attribute(self.file_id, key, str(value), group_id)

        return True

    def extract_metadata(self, metadata, group_name, parent_group_id):
        group_id = self.database.add_attribute_group(self.file_id, group_name, parent_group_id)
        for (key, value) in metadata.items():
            if isinstance(value, dict):
                self.extract_metadata(value, key, group_id)
            elif isinstance(value, tuple):
                self.database.add_attribute(self.file_id, key, self.tuple_to_string(value), group_id)
            else:
                self.database.add_attribute(self.file_id, key, str(value), group_id)

    def get_preview(self):
        with nio.dm.fileDM(".temp/"+self.file_id) as dm:
            data = dm.getDataset(0)["data"]
            data = np.array(data)
            data = np.rot90(data, 3)
            data = np.flip(data, 1)

            data = data - np.min(data)
            data = data / np.max(data)
            data = data * 255
            data = data.astype(np.uint8)

            img = Image.fromarray(data)
            img.save(".temp/"+self.file_id.split(".")[0] + ".png")


        url = "https://synopticprojectstorage.blob.core.windows.net/"
        token_credential = DefaultAzureCredential()
        blob_service_client = BlobServiceClient(account_url=url, credential=token_credential)
        container_client = blob_service_client.get_container_client(container="fyp-previews")
        with open(file=".temp/"+self.file_id.split(".")[0] + ".png", mode="rb") as data:
            container_client.upload_blob(name=self.file_id.split(".")[0] + ".png", data=data, overwrite=True)

        os.remove(".temp/"+self.file_id.split(".")[0] + ".png")
