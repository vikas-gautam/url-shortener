# Base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GOOS=linux CGO_ENABLED=0 go build -o gateway .

RUN chmod +x /app/gateway

#build tiny app
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/gateway /app

CMD [ "/app/gateway" ]
