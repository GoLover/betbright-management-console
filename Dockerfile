FROM golang:1.18-alpine AS build

RUN apk add git

RUN apk add librdkafka librdkafka-dev

RUN apk add gcc musl-dev

WORKDIR /tmp/go

COPY go.mod .
RUN git config --global http.sslVerify false

RUN GOPROXY=goproxy.io go mod download

COPY . .

RUN go build -tags musl -o ./out/out.o .

FROM alpine:latest

RUN apk add tzdata

COPY --from=build /tmp/go/out/out.o /app/out.o

CMD ["/app/out.o"]

