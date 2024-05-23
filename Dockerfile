############### BUILD STAGE ###############
FROM golang:1.22.3-alpine as builder
RUN apk update && apk upgrade && apk add --no-cache ca-certificates make

WORKDIR /app
ADD . /app
RUN make build

############### FINAL STAGE ###############
FROM scratch

COPY --from=builder /app/bin/captain-hindsight .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./captain-hindsight"]