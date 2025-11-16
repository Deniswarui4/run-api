#!/bin/bash

echo "🗄️  Setting up PostgreSQL Database with Docker"
echo "=============================================="

# Stop and remove existing container if it exists
echo "Cleaning up old containers..."
sudo docker stop event_ticketing_db 2>/dev/null || true
sudo docker rm event_ticketing_db 2>/dev/null || true

# Start PostgreSQL container
echo "Starting PostgreSQL container..."
sudo docker run -d \
  --name event_ticketing_db \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=event_ticketing \
  -p 5432:5432 \
  postgres:15-alpine

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
for i in {1..30}; do
    if sudo docker exec event_ticketing_db pg_isready -U postgres > /dev/null 2>&1; then
        echo "PostgreSQL is ready!"
        break
    fi
    echo "Waiting... ($i/30)"
    sleep 1
done

# Create test database
echo "Creating test database..."
sudo docker exec event_ticketing_db psql -U postgres -c "CREATE DATABASE event_ticketing_test;" 2>/dev/null || echo "Test database may already exist"

# Verify connection
echo "Verifying database connection..."
sudo docker exec event_ticketing_db psql -U postgres -c "SELECT version();" > /dev/null

if [ $? -eq 0 ]; then
    echo "✅ PostgreSQL is running successfully!"
    echo ""
    echo "Database Details:"
    echo "  Host: localhost"
    echo "  Port: 5432"
    echo "  Database: event_ticketing"
    echo "  Test Database: event_ticketing_test"
    echo "  User: postgres"
    echo "  Password: postgres"
    echo ""
    echo "Next steps:"
    echo "  1. Update your .env file with these credentials"
    echo "  2. Run: go run scripts/seed_admin.go"
    echo "  3. Run: go test ./... -v"
else
    echo "❌ Failed to connect to database"
    exit 1
fi
