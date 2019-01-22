#!/bin/bash
docker build -t spacelama/api:go-latest .
docker push spacelama/api:go-latest