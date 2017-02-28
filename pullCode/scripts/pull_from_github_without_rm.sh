#!/bin/sh

git clone https://$GITHUB_USER:$GITHUB_PWD@github.com/$REPO_URL $DEST
cd $DEST
git checkout $BRANCH
exit 0
