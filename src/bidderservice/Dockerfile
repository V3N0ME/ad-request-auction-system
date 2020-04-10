FROM golang:1.13 AS build-img
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o /bidderservice cmd/bidderservice/*.go

# Build app
FROM alpine:latest
WORKDIR /home

COPY --from=build-img /bidderservice .

# Start apps
ENV ENV=production
ENV TZ=Asia/Kolkata

CMD ./bidderservice