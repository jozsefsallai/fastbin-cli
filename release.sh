#!/bin/bash

if ! [ -x "$(command -v gox)" ]; then
  echo 'gox not found, please install it: https://github.com/mitchellh/gox'
  exit 1
fi

mkdir -p .release
rm -r .release/*
gox -os="linux darwin windows openbsd" -arch="386 amd64 arm arm64" -output=".release/{{.Dir}}_{{.OS}}_{{.Arch}}"
