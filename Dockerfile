FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o myapp cmd/main.go cmd/wire_gen.go

EXPOSE 8080

CMD ["./myapp"]