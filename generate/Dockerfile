FROM python:3.10-alpine
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . /app/generate
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["python -m generate"]