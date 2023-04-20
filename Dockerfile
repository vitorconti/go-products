FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git

WORKDIR /

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN wire
RUN go get github.com/google/wire/cmd/wire && \
    wire ./cmd && \
    go build -o go-products-api ./cmd/main.go

EXPOSE 8080

CMD ["./go-products-api"]