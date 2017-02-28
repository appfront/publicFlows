#!/bin/sh

cd /tmp
wget $URL
unzip $NAME.zip
mv $NAME/* $DEST
exit 0
