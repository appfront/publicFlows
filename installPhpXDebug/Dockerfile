# Set the base image to Ubuntu
FROM alpine:latest

# File Author / Maintainer
LABEL MAINTAINER Daniel Marzini - Appfront

COPY xdebug.ini /opt/
COPY app/configure /usr/bin/

RUN rm -rf /var/cache/apk/*
ENTRYPOINT /usr/bin/configure
