FROM       rayyildiz/java8:jre
MAINTAINER Ramazan AYYILDIZ <rayyildiz@gmail.com>

ENV  HBASE_VERSION=2.1.0

RUN apk add --progress --update unzip tar bash

RUN cd /tmp && wget http://artfiles.org/apache.org/hbase/${HBASE_VERSION}/hbase-${HBASE_VERSION}-bin.tar.gz && \
	tar -xzf hbase-${HBASE_VERSION}-bin.tar.gz 

RUN mkdir -p /usr/lib/hbase && \
    mv "/tmp/hbase-${HBASE_VERSION}" /usr/lib/hbase && \
    ln -s /usr/lib/hbase/hbase-${HBASE_VERSION} /usr/lib/hbase/hbase-latest 

RUN  apk del wget ca-certificates curl unzip tar && \
     rm -rf /tmp/* && \
     rm -rf /var/cache/apk/* && \
     rm -rf /usr/lib/hbase/hbase-${HBASE_VERSION}/docs

RUN chmod +x /usr/lib/hbase/hbase-latest/bin/start-hbase.sh && \
	chmod +x /usr/lib/hbase/hbase-latest/bin/stop-hbase.sh && \
	chmod +x /usr/lib/hbase/hbase-latest/bin/zookeepers.sh

EXPOSE 16000 16010 2888 3888 2181

CMD ["/usr/lib/hbase/hbase-latest/bin/start-hbase.sh"]