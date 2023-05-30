#!/bin/bash

echo "==> Running unit tests..."
go clean -testcache ./...
go test $(go list ./...) -p 1 --cover
