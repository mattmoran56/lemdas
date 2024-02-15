from fastapi import APIRouter
from pydantic import BaseModel

from processor import TiffScan

router: APIRouter = APIRouter()

class ProcessArgs(BaseModel):
    file_id: str

@router.post("/process/")
def read_root(args: ProcessArgs):
    processor = TiffScan(args.file_id)
    processor.download_files()
    processor.process()
    processor.finish_process()

    return True