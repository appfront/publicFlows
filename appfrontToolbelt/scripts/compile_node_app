#!/bin/sh

node --version

npm install -g bower
npm install -g grunt
npm install -g grunt-cli

grunt --version
bower --version

cd $DEST

npm install
bower install --allow-root | xargs echo
#npm test
grunt $BUILD_VERSION
mv $DEST/dist /opt
rm -rf /var/www/*
mv /opt/dist /var/www/$APP_NAME
