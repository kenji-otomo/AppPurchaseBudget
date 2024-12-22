#!/bin/bash

if [ -f .env.prd ]; then
  export $(cat .env.prd | xargs)
fi

docker build \
  --build-arg DNS="$DNS" \
  --build-arg VUE_URL="$VUE_URL" \
  -t gcr.io/kakinbo-445308/go-kakinbo-app .