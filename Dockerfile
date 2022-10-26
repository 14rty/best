FROM golang:1.19-alpine3.16
USER root
COPY ./api/* /api/
RUN apk update && go get github.com/lib/pq
WORKDIR /api/
EXPOSE 5432
ENTRYPOINT ["go", "./data-access.go"]