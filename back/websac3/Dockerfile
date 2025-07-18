FROM golang:1.23.7-alpine3.21 AS builder

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    echo "America/Bogota" > /etc/timezone

WORKDIR /app

COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4
RUN swag init --dir /app --output /app/docs --parseDependency --parseInternal

RUN go build -o /bin/anise github.com/JorgeGorrito/anise-with-gin/anise/command/client
RUN go build -o /bin/main .


FROM alpine:3.21

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    echo "America/Bogota" > /etc/timezone

WORKDIR /app

COPY --from=builder /bin/main /app/main
COPY --from=builder /bin/anise /app/anise
COPY --from=builder /app /app

COPY .env .
COPY .mail-credentials.json .
COPY .mail-token.json .

COPY adapter/out/message/lang/ adapter/out/message/lang/
COPY adapter/out/notification/template/templates adapter/out/notification/template/templates
COPY adapter/out/persistence/postgresql/seeders/seeds/ adapter/out/persistence/postgresql/seeders/seeds/

CMD ["./main"]
