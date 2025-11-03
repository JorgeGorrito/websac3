FROM golang:1.23.7-bookworm AS builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends tzdata ca-certificates && \
    ln -fs /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4
RUN swag init --dir /app --output /app/docs --parseDependency --parseInternal

RUN go build -o /bin/anise github.com/JorgeGorrito/anise-with-gin/anise/command/client
RUN go build -o /bin/main .


FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends tzdata ca-certificates wkhtmltopdf fontconfig fonts-dejavu openssl bash && \
    ln -fs /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /bin/main /app/main
COPY --from=builder /bin/anise /app/anise

COPY .env .
COPY ./.*.json .

COPY adapter/out/message/lang/ adapter/out/message/lang/
COPY adapter/out/notification/template/templates adapter/out/notification/template/templates
COPY adapter/out/persistence/postgresql/seeders/seeds/ adapter/out/persistence/postgresql/seeders/seeds/
COPY ensure-certs.sh /app/ensure-certs.sh

RUN sed -i 's/\r$//' /app/ensure-certs.sh && \
    chmod +x /app/ensure-certs.sh

ENTRYPOINT ["/bin/bash", "/app/ensure-certs.sh"]
CMD ["./main"]
