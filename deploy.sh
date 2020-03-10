#!/bin/bash
image=spacelama/web
# Uncomment for debug
# set -x

# User arguments
response=$1
message=$2
newTag=$3

# Get version from version.txt
majVersion=$(cat ./version.txt | cut -d'.' -f1)
minVersion=$(cat ./version.txt | cut -d'.' -f2)
echo current version : $majVersion.$minVersion

# Ask user for major or minor update
if [[ -z "$response" ]]; then
    echo [Git: 1/3] major or minor update ? [major/minor, default: minor]
    read response
else
    if [[ "$response" == "major" ]]; then
        majVersion=$((majVersion + 1))
        minVersion=0
    else
        minVersion=$((minVersion + 1))
    fi
fi

#  Ask user for commit message
if [[ -z "$message" ]]; then
    echo [Git: 2/3] To you to add commit message ? [default: $majVersion.$minVersion]
    read message
    if [[ -z "$message" ]]; then
        message='version: '$majVersion'.'$minVersion
    fi
fi
#  Ask user if he want to generate new tag
if [[ -z "$newTag" ]]; then
    echo [Git: 3/3] Do you want to generate new tag ? [y/n default: n]
    read newTag
    if [[ -z "$newTag" ]]; then
        newTag='n'
    fi
fi


echo $majVersion.$minVersion > version.txt
git add --all
git commit -m "$message"

if [[ "$newTag" == "y" ]]; then
    git tag -a "v$majVersion.$minVersion" -m "$message"
fi

git push

FILE=./Dockerfile
if [[ -f "$FILE" ]]; then
    echo "Dockerfile found on this folder ! "
    echo [Docker: 1/2] Do you want to build ? [y/n default: n]
    read build
    
    if [[ "$build" == "y" ]]; then
        echo [Docker: 2/2] Please choose branch to deploy $image ? [latest/dev] default: dev
        read branch
        
        if [[ $branch == "latest" ]]
        then
            tag="latest"
        else
            tag="dev"
        fi
        
        echo "Deploying $image:$tag on Docker.io"
        docker build -t $image:$tag .;
        docker push $image:$tag;
        echo "Done !"
    fi
fi

