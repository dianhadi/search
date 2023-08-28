FROM golang:alpine

RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache curl

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o binary ./cmd/consumer

# EXPOSE 8080

ENTRYPOINT ["./binary"]