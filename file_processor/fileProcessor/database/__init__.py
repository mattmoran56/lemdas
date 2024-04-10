import sqlalchemy
import os
import uuid

from sqlalchemy import select, update, insert

from database.models import File, FileAttribute, FileAttributeGroup


class FileDatabase:
    def __init__(self):
        self.db_username = os.environ.get("DB_USERNAME")
        self.db_password = os.environ.get("DB_PASSWORD")
        self.db_host = os.environ.get("DB_HOST")
        self.db_port = os.environ.get("DB_PORT")
        self.db_name = os.environ.get("DB_NAME")

        connection_string = f"mysql+mysqlconnector://{self.db_username}:{self.db_password}@{self.db_host}:{self.db_port}/{self.db_name}"
        engine = sqlalchemy.create_engine(connection_string)
        self.engine = engine

    def get_file_by_id(self, file_id):
        with self.engine.connect() as conn:
            stmt = select(File).where(File.id == file_id)
            return conn.execute(stmt).first()

    def get_file_by_name_dataset(self, file_name, dataset_id):
        with self.engine.connect() as conn:
            stmt = select(File).where(File.name == file_name, File.dataset_id == dataset_id)
            return conn.execute(stmt).first()

    def update_status(self, file_id, status):
        with self.engine.connect() as conn:
            stmt = update(File).where(File.id == file_id).values(status=status)
            conn.execute(stmt)
            conn.commit()

    def add_attribute(self, file_id, attribute_name, attribute_value, attribute_group_id):
        with self.engine.connect() as conn:
            id = str(uuid.uuid4())
            stmt = insert(FileAttribute).values(id=id, file_id=file_id, attribute_name=attribute_name, attribute_value=attribute_value, attribute_group_id=attribute_group_id)
            conn.execute(stmt)
            conn.commit()

    def add_attribute_group(self, file_id, attribute_group_name, parent_group_id=None):
        with self.engine.connect() as conn:
            id = str(uuid.uuid4())
            stmt = insert(FileAttributeGroup).values(id=id, file_id=file_id, attribute_group_name=attribute_group_name, parent_group_id=parent_group_id)
            conn.execute(stmt)
            conn.commit()

            return id