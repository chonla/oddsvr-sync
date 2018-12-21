FROM golang:1.11.3-alpine AS builder

WORKDIR /go/src/github.com/chonla/oddsvr-sync

COPY . .

RUN apk add --no-cache git \
    && go get ./... \
    && GOOS=linux GOARCH=amd64 go build -o oddsvr-sync



FROM alpine:latest

ARG ODDSVR_STRAVA_CLIENT_ID
ARG ODDSVR_STRAVA_CLIENT_SECRET
ARG ODDSVR_DB

ENV ODDSVR_STRAVA_CLIENT_ID=${ODDSVR_STRAVA_CLIENT_ID}
ENV ODDSVR_STRAVA_CLIENT_SECRET=${ODDSVR_STRAVA_CLIENT_SECRET}
ENV ODDSVR_DB=${ODDSVR_DB}

COPY --from=builder /go/src/github.com/chonla/oddsvr-sync/oddsvr-sync .
COPY --from=builder ./cron .

RUN cat ./cron >> /etc/crontabs/root

RUN ./oddsvr-sync

CMD ["crond", "-f", "-d", "8"]