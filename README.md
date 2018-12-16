# ODDS Virtual Run Data Synchronizing Tool

## Execution

```
ODDSVR_STRAVA_CLIENT_ID=<strava client id> ODDSVR_STRAVA_CLIENT_SECRET=<strava client secret> ODDSVR_DB=<mongo db connection> go run main.go
```

## Environment Variables

| Name | Description |
| - | - |
| ODDSVR_STRAVA_CLIENT_ID | Strava Client ID. See Strava Developer page for detail |
| ODDSVR_STRAVA_CLIENT_SECRET | Strava Client Secret. See Strava Developer page for detail |
| ODDSVR_DB | MongoDB Connection. `<username>:<password>@<mongodb-host>:<mongodb-port>` |

## Try with Docker

```
# Build image
docker build -t oddsvr_sync --build-arg ODDSVR_STRAVA_CLIENT_ID='<strava client id>' --build-arg ODDSVR_STRAVA_CLIENT_SECRET='<strava client secret>' --build-arg ODDSVR_DB='<mongo db connection>' .

# Run as a container
docker run -d --name oddsvr_sync -it oddsvr_sync

# Check log
docker logs `docker ps -aq -f name=oddsvr_sync`

# Stop the container
docker stop `docker ps -aq -f name=oddsvr_sync`

# Remove the container
docker rm `docker ps -aq -f name=oddsvr_sync`
```