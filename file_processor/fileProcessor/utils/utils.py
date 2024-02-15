from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient

def download_from_azure(file_id: str):
    account_url = "https://synopticprojectstorage.blob.core.windows.net/"
    default_credential = DefaultAzureCredential()
    blob_service_client = BlobServiceClient(account_url, credential=default_credential)
    blob_client = blob_service_client.get_blob_client(container="fyp-uploads", blob=file_id)
    with open(".temp/" + file_id, 'wb') as sample_blob:
        download_stream = blob_client.download_blob()
        sample_blob.write(download_stream.readall())