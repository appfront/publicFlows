#!/bin/sh

git clone https://$GITHUB_USER:$GITHUB_PWD@github.com/$REPO_URL $DEST
cd $DEST
git checkout $BRANCH
export LOG=`git log -1`
export MESSAGE="building "+ $REPO_URL +"\n"+$LOG
rm -rf .git
/opt/postToSlack
exit 0
