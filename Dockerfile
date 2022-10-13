FROM python:3.10-alpine3.16
USER root
COPY ./api/* /api/
RUN apk update && pip install --upgrade pip && apk add python3-dev gcc g++ python3-dev musl-dev && pip install -r api/req.txt
WORKDIR /api/
EXPOSE 5432
ENTRYPOINT ["python", "./test.py"]