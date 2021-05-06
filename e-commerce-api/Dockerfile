FROM golang:alpine as builder

WORKDIR /app
COPY ./ /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GODEBUG=madvdontneed=1 go build -a -tags netgo  -o app cmd/main.go

FROM alpine:latest as runner
WORKDIR /
COPY --from=builder /app .
EXPOSE 8081
ENTRYPOINT ["./app"]
