FROM golang:1.22.10-alpine3.21

ENV ROOT /devport

WORKDIR ${ROOT}

RUN apk update && apk add git

COPY . ${ROOT}

RUN go mod tidy

RUN go test ./test/**

RUN go install github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]
