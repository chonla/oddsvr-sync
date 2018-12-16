FROM golang:1.11.3-alpine

WORKDIR /go/src/github.com/chonla/oddsvr-sync

COPY . .

ARG ODDSVR_STRAVA_CLIENT_ID
ARG ODDSVR_STRAVA_CLIENT_SECRET
ARG ODDSVR_DB

ENV ODDSVR_STRAVA_CLIENT_ID=${ODDSVR_STRAVA_CLIENT_ID}
ENV ODDSVR_STRAVA_CLIENT_SECRET=${ODDSVR_STRAVA_CLIENT_SECRET}
ENV ODDSVR_DB=${ODDSVR_DB}

RUN apk add --no-cache git \
    && go get ./... \
    && go build \
    && cat ./cron >> /etc/crontabs/root

RUN /go/src/github.com/chonla/oddsvr-sync/oddsvr-sync

CMD ["crond", "-f", "-d", "8"]