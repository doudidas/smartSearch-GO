#!/bin/bash

# First we build the tool
cd ./initDB
go build -o ../build/initDB

# Build main app
cd ../
go build -o build/app 


# run in background
build/initDB &
build/app