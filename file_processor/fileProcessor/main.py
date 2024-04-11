import uvicorn
import os

from web import web

if __name__ == "__main__":
    uvicorn.run(web, port=8080, host="0.0.0.0")