FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux make output
FROM scratch
COPY --from=builder /build/output /app/
EXPOSE 7060
WORKDIR /app/output
CMD [ "bin/server", "-c", "configs/server.json" ]
