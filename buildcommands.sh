#!/usr/bin/env bash
set -e

mkdir -p build build/bin

GOOS=linux GOARCH=amd64 go build -o build/MatchZy-Webhook
GOOS=windows GOARCH=amd64 go build -o build/MatchZy-Webhook.exe

tar -cvzf build/bin/MatchZy-Webhook-Linux-amd64.tar.gz \
  build/MatchZy-Webhook build/config build/templates.json

zip -r build/bin/MatchZy-Webhook-Windows-amd64.zip \
  build/MatchZy-Webhook.exe build/config build/templates.json