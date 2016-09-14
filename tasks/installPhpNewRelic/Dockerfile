# Set the base image to Ubuntu
FROM alpine

# File Author / Maintainer
MAINTAINER Daniel Marzini

COPY newrelic.ini /opt/
COPY app/configure /usr/bin/

RUN rm -rf /var/cache/apk/*
ENTRYPOINT /usr/bin/configure
