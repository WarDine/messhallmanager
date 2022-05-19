FROM golang:1.16-alpine as base

FROM base as dev

ADD . /go/src/app
WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY .env ./

RUN go mod download
RUN go mod tidy
RUN go build -o ./recipe-manager-service

CMD [ "./recipe-manager-service" ]
