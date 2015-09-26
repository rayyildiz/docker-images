FROM       rayyildiz/java:jdk8
MAINTAINER Ramazan AYYILDIZ <rayyildiz@gmail.com>

ENV  ACTIVATOR_VERSION=1.3.6

RUN apk add --progress --update unzip bash 

RUN cd /tmp && wget http://downloads.typesafe.com/typesafe-activator/${ACTIVATOR_VERSION}/typesafe-activator-${ACTIVATOR_VERSION}.zip 

RUN cd /tmp && unzip typesafe-activator-${ACTIVATOR_VERSION}.zip 

RUN mkdir -p /usr/lib/activator && mv "/tmp/activator-dist-${ACTIVATOR_VERSION}"  \ 
	 "/usr/lib/activator/activator-${ACTIVATOR_VERSION}"

RUN ln -s /usr/lib/activator/activator-${ACTIVATOR_VERSION}/activator /usr/bin/activator

RUN  apk del wget ca-certificates curl unzip && \
     rm /tmp/* /var/cache/apk/*

EXPOSE 8888 8888

RUN mkdir /app
WORKDIR /app

CMD ["/usr/bin/activator", "ui", "-Dhttp.address=0.0.0.0"]
