# Set the base image to Ubuntu
FROM alpine

# File Author / Maintainer
MAINTAINER Daniel Marzini

RUN apk add --update ca-certificates

COPY app/exec /usr/bin/

RUN rm -rf /var/cache/apk/*
ENTRYPOINT /usr/bin/exec
