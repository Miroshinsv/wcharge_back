FROM golang:1.21.1-alpine3.18 AS builder

RUN apk update && \
    apk add bash ca-certificates git gcc g++ libc-dev binutils file

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /opt/application .

FROM alpine:3.18 AS production
RUN apk update && \
    apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY --from=builder /opt/application ./
# COPY --from=builder /opt/migration/* ./migration
# COPY --from=builder /opt/config/* ./config

CMD ["./application"]