#!/bin/bash

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Build the Go application
go build -o my-go-backend ./cmd/server

# Run the application
./my-go-backend &

# Optionally, you can add a command to wait for the application to start
# sleep 2

# Open the browser to the frontend URL (if applicable)
# open http://localhost:3000