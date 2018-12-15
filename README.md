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