FROM golang:1.24.5-alpine AS builder

WORKDIR /zerome

ADD go.mod go.sum ./
RUN go mod download

ADD . .

RUN apk add make git

RUN make build

FROM alpine:3.22.1

COPY --from=builder /zerome/zerome /usr/local/bin/zerome

ENTRYPOINT ["/usr/local/bin/zerome"]
