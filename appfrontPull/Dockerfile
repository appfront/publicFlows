# Set the base image to Ubuntu
FROM alpine

# File Author / Maintainer
MAINTAINER Daniel Marzini

RUN apk add --update ca-certificates python py-setuptools ruby 
RUN easy_install-2.7 s3cmd

COPY s3_config /root/.s3cfg
COPY pull /bin/
COPY push /bin/

RUN chmod +x /bin/pull
RUN chmod +x /bin/push


RUN rm -rf /var/cache/apk/* 

ENTRYPOINT /bin/sh
