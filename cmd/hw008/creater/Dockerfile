FROM golang:rc-alpine3.14

WORKDIR /app/token-creator

COPY . .

RUN go mod download github.com/dgrijalva/jwt-go

RUN go build -o ./out/main .

EXPOSE 6666

CMD ["./out/main"]
