FROM golang:1.16 AS builder

WORKDIR /go/src/app
COPY app/. .
ARG CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /go/app


FROM scratch
COPY --from=builder /go/app /go/app

ENTRYPOINT ["/go/app"]