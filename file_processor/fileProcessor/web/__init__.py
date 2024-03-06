from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse

import os
import traceback
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

@web.middleware("http")
async def catch_exceptions(request: Request, call_next):
    try:
        return await call_next(request)
    except Exception as e:
        # Log the exception for debugging purposes
        traceback.print_exc()

        # Construct the error response with stack trace
        error_detail = {
            "error": "Internal Server Error",
            "detail": str(e),
            "stacktrace": traceback.format_exc()
        }

        # Encode the error response to JSON
        error_response = jsonable_encoder(error_detail)

        # Return the JSON response with status code 500
        return JSONResponse(status_code=500, content=error_response)
@web.on_event("startup")
async def startup_event():
    # TODO: Connect to Azure
    print("Connect to azure")

@web.get("/")
def read_root():
    return {"message": "Service started updated"}


web.include_router(router)
