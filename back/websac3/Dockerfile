FROM golang:1.23.7-alpine3.21 AS builder

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    echo "America/Bogota" > /etc/timezone

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init

RUN go build -o /bin/anise github.com/JorgeGorrito/anise-with-gin/anise/command/client && \
    go build -o /bin/main .

FROM alpine:3.21

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    echo "America/Bogota" > /etc/timezone

WORKDIR /app

COPY --from=builder /bin/main /app/main
COPY --from=builder /bin/anise /app/anise

COPY .env .
COPY .mail-credentials.json .
COPY .mail-token.json .

COPY adapter/out/message/lang/ adapter/out/message/lang/
COPY adapter/out/notification/template/templates adapter/out/notification/template/templates
COPY adapter/out/persistence/postgresql/seeders/seeds/ adapter/out/persistence/postgresql/seeders/seeds/

CMD ["./main"]
