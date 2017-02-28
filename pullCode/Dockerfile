# Set the base image to Ubuntu
FROM alpine

# File Author / Maintainer
MAINTAINER Daniel Marzini

RUN apk add --update ca-certificates git bash
RUN mkdir /opt

COPY app/postToSlack /opt
COPY scripts /opt/scripts
RUN chmod +x /opt/postToSlack

RUN rm -rf /var/cache/apk/*
ENTRYPOINT /bin/sh
