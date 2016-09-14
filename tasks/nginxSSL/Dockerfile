# Set the base image to Ubuntu
FROM alpine

# File Author / Maintainer
MAINTAINER Daniel Marzini

RUN apk add --update ca-certificates nginx

COPY nginx.conf /etc/nginx/
COPY default.conf /etc/nginx/conf.d/

RUN rm -rf /var/cache/apk/*
ENTRYPOINT /usr/sbin/nginx -c /etc/nginx/nginx.conf
