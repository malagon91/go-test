FROM golang:1.16-alpine AS builder

WORKDIR /go/src/api-bootstrap-echo

COPY . .

RUN mkdir cfg; mv ./config/*.json ./cfg/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o api-bootstrap-echo .

FROM alpine:3.13

ARG version

ENV VERSION ${version}

WORKDIR /go/src/api-bootstrap-echo

COPY --from=builder /go/src/api-bootstrap-echo/cfg ./config

COPY --from=builder /go/src/api-bootstrap-echo/api-bootstrap-echo ./api-bootstrap-echo

EXPOSE 4001

CMD ["./api-bootstrap-echo"]
