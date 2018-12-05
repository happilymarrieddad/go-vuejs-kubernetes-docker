#!/bin/bash

if [ -z ${SHA+x} ]; then SHA=$(git rev-parse HEAD); fi

docker build -t 0xsegfault/echo1:lastest-dev -f ./echoservers/echo1/Dockerfile.dev ./echoservers/echo1
docker build -t 0xsegfault/echo1:$SHA -f ./echoservers/echo1/Dockerfile.dev ./echoservers/echo1

docker push 0xsegfault/echo1:lastest-dev
docker push 0xsegfault/echo1:$SHA
