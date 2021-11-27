# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
ENV DB postgres
WORKDIR /application
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
ADD /app ./app
EXPOSE 9000
RUN go build -o /docker-url-shortener
CMD ["sh", "-c", "/docker-url-shortener --db=$DB"]
