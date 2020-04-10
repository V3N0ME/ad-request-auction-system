FROM golang:1.13 AS build-img
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o /auctionservice cmd/auctionservice/*.go

# Build app
FROM alpine:latest
WORKDIR /home

COPY --from=build-img /auctionservice .

ENV ENV=production
ENV TZ=Asia/Kolkata
ENV PORT=8080

#ENV GIN_MODE=release

EXPOSE 8080

CMD ./auctionservice