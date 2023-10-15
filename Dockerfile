# syntax=docker/dockerfile:1.4
FROM golang:1.21-alpine3.18 AS builder

ARG VERSION

WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags="-X 'github.com/jjuarez/simple-prober/cmd.Version=${VERSION}'" -o dist/simple-prober main.go


FROM alpine:3.18.4 AS runtime

ARG UID=1001

COPY --from=builder --chown=1001:1001 /build/dist/simple-prober /usr/local/bin/simple-prober
WORKDIR /app
RUN mkdir -p config
USER ${UID}
VOLUME config/endpoints.yaml
CMD [ "simple-prober", "check", "--config", "config/endpoints.yaml", "--timeout", "5", "--loglevel", "debug" ]
