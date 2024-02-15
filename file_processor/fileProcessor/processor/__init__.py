import os
from utils.utils import download_from_azure

from database import FileDatabase

class TiffScan:
    def __init__(self, file_id):
        self.database = FileDatabase()
        self.halted = False

        self.file_id = file_id
        file = self.database.get_file_by_id(file_id)
        self.txt_file_name = file.name.replace(".tif", ".txt")
        print("Txt file name: " + self.txt_file_name)
        txt_file = self.database.get_file_by_name_dataset(self.txt_file_name, file.dataset_id)
        if txt_file is None:
            print("No txt file found")
            self.database.update_status(self.file_id, "awaitingtxt")
            self.halted = True
        else:
            print("Txt file found: " + txt_file.id)
            self.txt_file_id = txt_file.id


    def download_files(self):
        if self.halted:
            return False
        download_from_azure(self.file_id)
        download_from_azure(self.txt_file_id)

        print("Files downloaded")


    def process(self):
        if self.halted:
            return False
        self.database.update_status(self.file_id, "processing")
        self.database.update_status(self.txt_file_id, "processing")

        print("updated statuses")

        with open(".temp/"+self.txt_file_id, 'r') as file:
            for raw_line in file:
                line = raw_line.strip()
                print("Line: " + line)
                if "=" not in line:
                    continue
                if "   " in line:
                    primary_attribute, child_attributes = line.split("=", 1)
                    print("Multiple attributes")
                    attributes = child_attributes.split("   ")
                    print(attributes)
                    for attribute in attributes:
                        print("Attribute: " + attribute)
                        key, value = attribute.split("=")
                        if value != "":
                            print("Adding attribute: |" + primary_attribute + "-" + key + "| : |" + value + "|")
                            self.database.add_attribute(self.file_id, primary_attribute + "-" + key, value)
                    continue
                key, value = line.split("=")
                if value != "" or len(value) > 0:
                    print("Adding attribute: |" + key + "| : |" + value + "|")
                    self.database.add_attribute(self.file_id, key, value)

        self.database.update_status(self.file_id, "processed")
        self.database.update_status(self.txt_file_id, "support_processed")



    def finish_process(self):
        os.remove(".temp/"+self.file_id)
        os.remove(".temp/"+self.txt_file_id)
        return True
