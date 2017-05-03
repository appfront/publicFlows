# Set the base image to Ubuntu
FROM alpine

# File Author / Maintainer
MAINTAINER Daniel Marzini

RUN apk --update add nginx supervisor
COPY nginx.conf /etc/nginx/nginx.conf
COPY supervisord.ini /etc/supervisor.d/supervisord.ini

RUN rm -rf /var/cache/apk/*
EXPOSE 80
CMD ["/usr/bin/supervisord -n"]
