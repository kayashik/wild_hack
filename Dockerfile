# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /docker-gs-wild-hack
EXPOSE 8080
CMD [ "/docker-gs-wild-hack" ]
