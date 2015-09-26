FROM       alpine

RUN apk update --progress && \
 	apk upgrade --progress

# Install cURL
RUN apk add --progress --update bash curl ca-certificates tar

RUN rm /var/cache/apk/*