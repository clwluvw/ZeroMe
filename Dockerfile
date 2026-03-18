FROM golang:1.26-alpine AS builder

WORKDIR /zerome

ADD go.mod go.sum ./
RUN go mod download

ADD . .

RUN apk add make git

RUN make build

FROM alpine:latest

COPY --from=builder /zerome/zerome /usr/local/bin/zerome

ENTRYPOINT ["/usr/local/bin/zerome"]
