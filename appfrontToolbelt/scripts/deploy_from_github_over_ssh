#!/bin/sh

mkdir -p /root/.ssh
cd /root/.ssh/
wget $RSA_KEY_URL
mv $RSA_KEY_NAME id_rsa
chmod 0600 id_rsa
echo 'StrictHostKeyChecking no' > /root/.ssh/config
echo 'UserKnownHostsFile=/dev/null' >> /root/.ssh/config

git clone $REPO_URL $DEST
cd $DEST
git checkout $BRANCH
rm -rf .git
