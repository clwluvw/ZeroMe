FROM golang:1.22.5-alpine as builder

WORKDIR /zerome

ADD go.mod go.sum ./
RUN go mod download

ADD . .

RUN apk add make git

RUN make build

FROM alpine:3.20.1

COPY --from=builder /zerome/zerome /usr/local/bin/zerome

ENTRYPOINT ["/usr/local/bin/zerome"]
