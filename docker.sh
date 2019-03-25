#!/bin/sh
# Git adn dep are mendatory to get go dependancies
 apk update && apk add --no-cache git dep 
# Get dependancies with dep
dep ensure
# Build a scratch version of the code
# GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main
go build