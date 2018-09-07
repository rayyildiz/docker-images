FROM       rayyildiz/base
MAINTAINER Ramazan AYYILDIZ <rayyildiz@gmail.com>

# Java Version
ENV  JAVA_VERSION=8 \
     JAVA_UPDATE=101 \
     JAVA_BUILD=13  \
     JAVA_HOME=/usr/lib/jvm/current-java

# JDK: http://download.oracle.com/otn-pub/java/jdk/8u60-b27/jdk-8u60-linux-x64.tar.gz
# JRE: http://download.oracle.com/otn-pub/java/jdk/8u60-b27/jre-8u60-linux-x64.tar.gz
# http://download.oracle.com/otn-pub/java/jdk/8u101-b13/jdk-8u101-linux-x64.tar.gz

# JRE
RUN cd /tmp && wget --header "Cookie: oraclelicense=accept-securebackup-cookie;" \
        "http://download.oracle.com/otn-pub/java/jdk/${JAVA_VERSION}u${JAVA_UPDATE}-b${JAVA_BUILD}/jre-${JAVA_VERSION}u${JAVA_UPDATE}-linux-x64.tar.gz" && \
     tar xzf "jre-${JAVA_VERSION}u${JAVA_UPDATE}-linux-x64.tar.gz" && \
     mkdir -p /usr/lib/jvm && mv "/tmp/jre1.${JAVA_VERSION}.0_${JAVA_UPDATE}" "/usr/lib/jvm/java-${JAVA_VERSION}-oracle"  && \
     ln -s "java-${JAVA_VERSION}-oracle" $JAVA_HOME && \
     ln -s $JAVA_HOME/bin/java /usr/bin/java && \
     rm -rf $JAVA_HOME/*.txt && \
     rm -rf $JAVA_HOME/*.html && \
     apk del wget ca-certificates curl && \
     rm /tmp/* /var/cache/apk/*

# JDK
# RUN cd /tmp && wget --header "Cookie: oraclelicense=accept-securebackup-cookie;" \
#       "http://download.oracle.com/otn-pub/java/jdk/${JAVA_VERSION}u${JAVA_UPDATE}-b${JAVA_BUILD}/jdk-${JAVA_VERSION}u${JAVA_UPDATE}-linux-x64.tar.gz" && \
#    tar xzf "jdk-${JAVA_VERSION}u${JAVA_UPDATE}-linux-x64.tar.gz" && \
#    mkdir -p /usr/lib/jvm && mv "/tmp/jdk1.${JAVA_VERSION}.0_${JAVA_UPDATE}" "/usr/lib/jvm/java-${JAVA_VERSION}-oracle"  && \
#    ln -s "java-${JAVA_VERSION}-oracle" $JAVA_HOME && \
#    ln -s $JAVA_HOME/bin/java /usr/bin/java && \
#    rm -rf $JAVA_HOME/*.txt && \
#    rm -rf $JAVA_HOME/*.html && \
#    apk del wget ca-certificates curl && \
#    rm /tmp/* /var/cache/apk/*
