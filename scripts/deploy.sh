#!/bin/bash
echo Please choose branch to deploy ? [latest/dev] default: dev
read branch

if [ "$branch" == "latest" ]
then
    tag="latest"
else
    tag="dev"
fi

docker build -t spacelama/api-go:$tag .
docker push spacelama/api-go:$tag