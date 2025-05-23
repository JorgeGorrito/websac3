FROM golang:1.23.7-alpine3.21

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Bogota /etc/localtime && \
    echo "America/Bogota" > /etc/timezone

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

RUN go build -o anise github.com/JorgeGorrito/anise-with-gin/anise/command/client 
RUN go build -o main .

CMD ["./main"]