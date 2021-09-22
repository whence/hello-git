#!/usr/bin/env bash
set -euo pipefail

cd /app
go mod edit -replace github.com/libgit2/git2go/v32=../git2go
go mod tidy
go build -tags static -o server main.go
