#!/usr/bin/env bash


# Get the git commit
GIT_COMMIT=$(git rev-parse --short HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_IMPORT=github.com/hashicorp/consul/version
GIT_DATE=$(./build-support/scripts/build-date.sh)
GOLDFLAGS=" -X $GIT_IMPORT.GitCommit=$GIT_COMMIT$GIT_DIRTY -X $GIT_IMPORT.BuildDate=$GIT_DATE "

go build -ldflags "$GOLDFLAGS" -o dist/ .

docker build -t consul-dev -f ./build-support/docker/Consul-Dev-windows.local.dockerfile .