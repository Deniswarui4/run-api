#!/bin/bash

# Script to restart the API server with the latest code

echo "🔄 Restarting Event Ticketing API..."

# Kill any existing Go processes for this API
pkill -f "go run.*cmd/api/main.go" || true
pkill -f "event-ticketing-api" || true

# Wait a moment for processes to stop
sleep 1

# Start the API server
echo "🚀 Starting API server..."
cd "$(dirname "$0")"
go run cmd/api/main.go
