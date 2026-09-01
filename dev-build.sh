#!/usr/bin/env bash

# Build and output directly to GOPATH bin as 'xdoit'
echo "Building xdoit system-wide..."
go build -o ~/go/bin/xdoit ./cmd/todo

if [ $? -eq 0 ]; then
    echo "Successfully updated 'xdoit' in ~/go/bin/xdoit"
else
    echo "Build failed!"
fi
