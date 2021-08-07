FROM golang:rc-alpine3.14

WORKDIR /app/token-storage

COPY . .

RUN go build -o ./out/main .

EXPOSE 7777

CMD ["./out/main"]
