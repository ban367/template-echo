FROM golang:1.13-alpine as build

WORKDIR /go/src
COPY . .

RUN apk add --update --no-cache git \
  && go build -o app \
  && GO111MODULE=off go get github.com/oxequa/realize

FROM alpine

WORKDIR /app
COPY --from=build /go/src/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /src/app

CMD ["./app"]
