# syntax=docker/dockerfile:1.4
FROM golang:1.19-alpine3.17 AS builder

ARG VERSION

WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags="-X 'github.com/jjuarez/simple-prober/cmd.Version=${VERSION}'" -o dist/simple-prober main.go


FROM alpine:3.17.1 AS runtime

WORKDIR /app
RUN mkdir -p config
USER 1001
COPY --from=builder --chown=1001:1001 /build/dist/simple-prober ./simple-prober
VOLUME /app/config
ENTRYPOINT [ "/app/simple-prober" ]
CMD [ "check", "--config", "/app/config/endpoints.yaml" ]
