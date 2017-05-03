FROM alpine:latest

RUN apk --update add wget curl unzip git
ADD scripts/* /opt/appfront/
RUN chmod +x /opt/appfront/*
RUN mv /opt/appfront/* /usr/local/bin/
RUN rm -rf /var/cache/apk/*
WORKDIR /opt/appfront
ENTRYPOINT []
