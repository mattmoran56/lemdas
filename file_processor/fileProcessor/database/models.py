from sqlalchemy import Column, String, ForeignKey, DateTime, Boolean
from sqlalchemy.sql import func
from sqlalchemy.orm import declared_attr
from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()


class BaseAbstract(object):
    @declared_attr
    def __tablename__(cls):
        return cls.__name__.lower()

    __table_args__ = {'mysql_engine': 'InnoDB'}
    id = Column(String, primary_key=True)
    created_at = Column(DateTime, default=func.now())
    updated_at = Column(DateTime, default=func.now(), onupdate=func.now())


class File(Base, BaseAbstract):
    __tablename__ = "files"

    name = Column(String)
    owner_id = Column(String)
    status = Column(String)
    dataset_id = Column(String)
    is_public = Column(Boolean, default=False)


class FileAttribute(Base, BaseAbstract):
    __tablename__ = "file_attributes"

    file_id = Column(String, ForeignKey("file.id"))
    attribute_name = Column(String)
    attribute_value = Column(String)
