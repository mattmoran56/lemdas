from utils.utils import download_from_azure
import os

class Scan:
    def download_files(self):
        download_from_azure(self.file_id)
        if self.halted:
            return False
        if hasattr(self, 'txt_file_id'):
            download_from_azure(self.txt_file_id)

        return True

    def finish_process(self):
        os.remove(".temp/" + self.file_id)
        if self.halted:
            return False
        if hasattr(self, 'txt_file_id'):
            os.remove(".temp/"+self.txt_file_id)
        return True