FROM golang:1.14.0-alpine3.11 AS builder

RUN go version
RUN apk add git

COPY . /github.com/fishmanDK/jwt-auth
WORKDIR /github.com/fishmanDK/jwt-auth

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app/cmd/jwt-auth/grpc/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/fishmanDK/jwt-auth/.bin/app .
COPY --from=0 /github.com/fishmanDK/jwt-auth/configs ./configs/

EXPOSE 5001

CMD ["./app"]