#!/usr/bin/env sh
set -eu

: "${PORT:=6060}"
: "${GO_VERSION:=latest}"

docker run --rm -it \
  -p "${PORT}:8080" \
  -v "$(pwd)":/src \
  -w /src \
  "golang:${GO_VERSION}" sh -c 'go install golang.org/x/pkgsite/cmd/pkgsite@latest && pkgsite -http=:8080'