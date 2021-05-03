FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build -o api api.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo
ENV PORT=5000

COPY --from=builder /app/api .
COPY config.Development.yaml .
COPY config.Staging.yaml .
COPY config.Production.yaml .

EXPOSE 5000

CMD [ "/app/api" ]