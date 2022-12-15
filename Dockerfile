FROM golang:1.19-alpine3.16
USER root
COPY ./api/* /api/
RUN apk update 
WORKDIR /api/
RUN go mod tidy
EXPOSE 5432
RUN go build -o data-access.exe .
ENTRYPOINT ["./data-access.exe"]
