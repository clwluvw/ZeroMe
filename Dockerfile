FROM golang:1.24.3-alpine AS builder

WORKDIR /zerome

ADD go.mod go.sum ./
RUN go mod download

ADD . .

RUN apk add make git

RUN make build

FROM alpine:3.21.3

COPY --from=builder /zerome/zerome /usr/local/bin/zerome

ENTRYPOINT ["/usr/local/bin/zerome"]
