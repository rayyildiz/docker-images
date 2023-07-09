FROM       alpine

RUN apk update --progress && \
 	apk upgrade --progress

# Install cURL
RUN apk add --progress --update bash ca-certificates

RUN rm /var/cache/apk/*
