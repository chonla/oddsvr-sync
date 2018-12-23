FROM golang:1.11.4-alpine AS builder

WORKDIR /go/src/github.com/chonla/oddsvr-sync

COPY . .

RUN apk add --no-cache git \
    && go get ./... \
    && GOOS=linux GOARCH=amd64 go build -o oddsvr-sync



FROM alpine:latest

WORKDIR /app

ARG ODDSVR_STRAVA_CLIENT_ID
ARG ODDSVR_STRAVA_CLIENT_SECRET
ARG ODDSVR_DB

ENV ODDSVR_STRAVA_CLIENT_ID=${ODDSVR_STRAVA_CLIENT_ID}
ENV ODDSVR_STRAVA_CLIENT_SECRET=${ODDSVR_STRAVA_CLIENT_SECRET}
ENV ODDSVR_DB=${ODDSVR_DB}

COPY --from=builder /go/src/github.com/chonla/oddsvr-sync/oddsvr-sync .
COPY --from=builder /go/src/github.com/chonla/oddsvr-sync/cron .

RUN cat /app/cron >> /etc/crontabs/root

CMD ["crond", "-f", "-d", "8"]