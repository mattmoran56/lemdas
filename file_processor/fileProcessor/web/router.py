from fastapi import APIRouter
from pydantic import BaseModel

from processor import Processor

router: APIRouter = APIRouter()

class ProcessArgs(BaseModel):
    file_id: str


@router.post("/process/")
def read_root(args: ProcessArgs):
    processor = Processor(args.file_id)
    processor.process()
