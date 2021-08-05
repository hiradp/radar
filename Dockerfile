FROM golang:1.16-alpine as builder

WORKDIR /radar

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/radar ./cmd/radar/main.go

FROM alpine

WORKDIR /radar

COPY --from=builder /radar/bin/radar /bin/radar
COPY assets ./assets

ENTRYPOINT ["/bin/radar"]
