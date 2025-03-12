#!/bin/sh

git filter-branch --env-filter '
if [ "$GIT_COMMITTER_EMAIL" != "$GIT_AUTHOR_EMAIL" ]; then
  export GIT_COMMITTER_NAME="$GIT_AUTHOR_NAME"
  export GIT_COMMITTER_EMAIL="$GIT_AUTHOR_EMAIL"
  export GIT_COMMITTER_DATE="$GIT_AUTHOR_DATE"
fi
' --tag-name-filter cat -- --branches --tags -f
