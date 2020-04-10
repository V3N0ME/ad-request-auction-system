# Ad Request Auction System

An auction system which can auction bids and select the winning bid such that it always responds before a specific time interval.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/8f742c30e73752fb6969)

## Running the project

```bash
$ docker-compose up --scale bidderservice=3
```

## Development

### Auction Service

```bash
$ cd src/auctionservice
$ go run cmd/auctionservice/main.go
```

### Bidder Service

```bash
$ cd src/bidderservice
$ go run cmd/bidderservice/main.go
```
