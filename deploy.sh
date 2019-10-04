#!/bin/bash

branchName = git branch | grep \* | cut -d ' ' -f2

if [ "$branch" == "master" ]
then
    tag="latest"
else
    tag=$branch
fi

docker build -t spacelama/api-go:$tag .
docker push spacelama/api-go:$tag