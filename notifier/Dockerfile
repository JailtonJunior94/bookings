FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o notifier notifier/notifier.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/notifier .
COPY ["notifier/appsettings.Development.json","."]
COPY ["notifier/appsettings.Production.json","."]
COPY ["notifier/appsettings.Staging.json","."]

CMD [ "/app/notifier" ]