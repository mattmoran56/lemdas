from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

import os
from pathlib import Path

from web.router import router

web: FastAPI = FastAPI()

origins = [
    "*",
]

web.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@web.on_event("startup")
async def startup_event():
    # TODO: Connect to Azure
    print("Connect to azure")

@web.get("/")
def read_root():
    return {"message": "Service started updated"}


web.include_router(router)
