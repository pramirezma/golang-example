FROM golang:1.15.6-alpine3.12 as build

WORKDIR /go/src/app

COPY . . 

RUN go build -o hello

FROM alpine:3.12

ENV PORT=3000

COPY --from=build /go/src/app/hello /usr/local/bin/app 

ENTRYPOINT ["/usr/local/bin/app"]