FROM       alpine
MAINTAINER Ramazan AYYILDIZ <rayyildiz@gmail.com>

# Install cURL
RUN apk add --update wget curl ca-certificates tar

RUN curl -Ls https://circle-artifacts.com/gh/andyshinn/alpine-pkg-glibc/6/artifacts/0/home/ubuntu/alpine-pkg-glibc/packages/x86_64/glibc-2.21-r2.apk > /tmp/glibc-2.21-r2.apk 

RUN apk add --allow-untrusted /tmp/glibc-2.21-r2.apk
