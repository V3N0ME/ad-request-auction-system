FROM golang:1.13 AS build-img
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o /bidderservice cmd/bidderservice/*.go

# Build app
FROM alpine:latest
WORKDIR /home

COPY --from=build-img /bidderservice .

ENV ENV=production
ENV TZ=Asia/Kolkata
ENV PORT=3000

ENV GIN_MODE=release

CMD ./bidderservice