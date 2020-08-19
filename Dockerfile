FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -ldflags "-X 'main.AppVersion=`sh scripts/version.sh`'" cmd/server/main.go
FROM scratch
COPY --from=builder /build/main /app/server
COPY --from=builder /build/configs/server/server.json  /app/
EXPOSE 7060
WORKDIR /app
CMD [ "./server", "-c", "./server.json" ]
