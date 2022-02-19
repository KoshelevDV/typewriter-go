FROM golang:1.17.5-alpine3.15 as base
WORKDIR /app
COPY . ./
RUN go build -o typeWriter ./cmd/main.go

FROM alpine:latest
COPY --from=base /app/typeWriter /usr/bin
RUN chmod +x /usr/bin/typeWriter
ENTRYPOINT [ "typeWriter" ]