FROM alpine as builder

RUN apk add --no-cache ca-certificates

FROM scratch
ENTRYPOINT []
WORKDIR /apps

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
