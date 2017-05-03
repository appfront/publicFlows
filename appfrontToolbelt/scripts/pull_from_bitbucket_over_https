#!/bin/sh

git clone https://$BITBUCKET_USER:$BITBUCKET_PWD@bitbucket.org/$REPO_URL $DEST
cd $DEST
git checkout $BRANCH
rm -rf .git
exit 0
