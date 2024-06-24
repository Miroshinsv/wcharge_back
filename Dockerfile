FROM golang:1.21.4-alpine3.18 AS builder

RUN apk update && \
    apk add bash ca-certificates git gcc g++ libc-dev binutils file

WORKDIR /opt

# COPY go.mod go.sum ./
COPY . .
RUN go mod download && go mod verify
RUN go build -tags migrate -o /opt/wcharge_back_migrate
RUN go build -o /opt/wcharge_back

FROM alpine:3.18 AS production
RUN apk update && \
    apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY --from=builder /opt/wcharge_back ./
COPY --from=builder /opt/wcharge_back_migrate ./
COPY ./migration/* ./migration/
COPY --from=builder /opt/config/* ./config

COPY run.sh .
RUN chmod +x run.sh
RUN echo 0 > migrate_flag

CMD ["./run.sh"]