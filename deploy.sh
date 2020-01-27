#!/bin/bash

# Get trace for debug
#  set -x

# User arguments
response=$1
message=$2
newTag=$3

#  Init version.txt file
if [ ! -f "./version.txt" ]; then
    echo "Init version File"
    echo "0.0" > ./version.txt
fi

# Get major/minor versions from version.txt
majVersion=$(cat ./version.txt | cut -d'.' -f1)
minVersion=$(cat ./version.txt | cut -d'.' -f2)
echo current version : $majVersion.$minVersion

# Ask user for major or minor update
if [[ -z "$response" ]]; then
    echo major or minor update ? [major/minor, default: minor]
    read response
    if [[ "$response" == "major" ]]; then
        majVersion=$((majVersion + 1))
        minVersion=0
    else
        minVersion=$((minVersion + 1))
    fi
fi


#  Ask user for commit message
if [[ -z "$message" ]]; then
    echo To you to add commit message ? [default: $majVersion.$minVersion]
    read message
    if [[ -z "$message" ]]; then
        message='version: '$majVersion'.'$minVersion
    fi
fi
#  Ask user if he want to generate new tag
if [[ -z "$newTag" ]]; then
    echo To you want to generate new tag ? [y/n default: n]
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