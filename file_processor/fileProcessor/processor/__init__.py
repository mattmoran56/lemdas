import database
from processor.tif_scan import TifScan


class Processor:
    def __init__(self, file_id):
        self.file_id = file_id
        self.file = database.FileDatabase().get_file_by_id(file_id)
        self.processing_started = False

        if self.file.status != "uploaded":
            print("File has already been processed, or is currently being processed")
            self.processing_started = True
            return
        self.file_type = file_id.split(".")[-1]

    def process(self):
        if self.processing_started:
            return

        database.FileDatabase().update_status(self.file_id, "processing")

        if self.file_type == "tif":
            processor = TifScan(self.file_id)
            processor.download_files()
            processor.process()
            processor.get_preview()
            processor.finish_process()

        elif self.file_type == "txt":
            tif_file_name = self.file.name.split(".")[0] + ".tif"
            tif_file = database.FileDatabase().get_file_by_name_dataset(tif_file_name, self.file.dataset_id)
            if tif_file is not None and tif_file.status == "awaitingtxt":
                processor = Processor(tif_file.id)
                processor.process()
            else:
                database.FileDatabase().update_status(self.file_id, "uploaded")
                return
        else:
            print("File type not supported")

        database.FileDatabase().update_status(self.file_id, "processed")
