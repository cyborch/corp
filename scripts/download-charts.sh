#!/bin/bash

cd "$(dirname "$0")/../docs/pages/charts"

for f in `grep tgz index.yaml | sed 's/ \- //g'` ; do \
  curl --silent -f -O $f || true
done
