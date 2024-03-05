FROM python:latest

WORKDIR /app

COPY .temp .
WORKDIR /app/file_processor/fileProcessor

ENV PATH="/root/.local/bin:$PATH"

RUN pip install --upgrade pip setuptools wheel

RUN pip install poetry==1.6.1
# RUN pip inject poetry setuptools==68.1.2

RUN poetry self add setuptools

RUN poetry config virtualenvs.in-project true
RUN poetry install --no-dev --no-ansi

EXPOSE 8080:8080

CMD poetry run python src/main.py