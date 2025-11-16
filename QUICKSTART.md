# Quick Start Guide

Get your Event Ticketing API up and running in 5 minutes!

## Prerequisites

- Go 1.21+
- PostgreSQL 12+
- libvips (for image processing)

### Install libvips

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install libvips-dev
```

**macOS:**
```bash
brew install vips
```

**Fedora/RHEL:**
```bash
sudo dnf install vips-devel
```

## Setup Steps

### 1. Clone and Navigate
```bash
cd event-ticketing-go-api
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Set Up Database
```bash
# Create PostgreSQL database
createdb event_ticketing

# Or using psql
psql -U postgres -c "CREATE DATABASE event_ticketing;"
```

### 4. Configure Environment
```bash
# Copy example environment file
cp .env.example .env

# Edit .env with your settings
nano .env
```

**Minimum required settings:**
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=event_ticketing

JWT_SECRET=your-super-secret-jwt-key-change-this
```

### 5. Seed Database
```bash
# Create admin and test users
go run scripts/seed_admin.go
```

This creates:
- **Admin**: admin@eventtickets.com / Admin@123
- **Moderator**: moderator@eventtickets.com / Moderator@123
- **Organizer**: organizer@eventtickets.com / Organizer@123
- **Attendee**: attendee@eventtickets.com / Attendee@123

### 6. Run the API
```bash
# Using go run
go run cmd/api/main.go

# Or build and run
make build
./bin/event-ticketing-api

# Or with hot reload (requires air)
make dev
```

The API will be available at: **http://localhost:8080**

### 7. Test the API
```bash
# Health check
curl http://localhost:8080/health

# Run test suite
./scripts/test_api.sh

# Or test login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@eventtickets.com",
    "password": "Admin@123"
  }'
```

## Using Docker (Alternative)

### Quick Start with Docker Compose
```bash
# Start all services (API + PostgreSQL)
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down
```

### Seed Database in Docker
```bash
# Run seed script in container
docker-compose exec api go run scripts/seed_admin.go
```

## Common Commands

```bash
# Build the application
make build

# Run tests
make test

# Run with coverage
make test-coverage

# Format code
make format

# View all commands
make help
```

## Next Steps

1. **Login** with one of the seeded accounts
2. **Create an event** as organizer
3. **Submit for review** 
4. **Approve event** as moderator
5. **Publish event** as organizer
6. **Purchase tickets** as attendee

## API Documentation

- Full API docs: [API_DOCUMENTATION.md](API_DOCUMENTATION.md)
- README: [README.md](README.md)

## Configuration for Production

### Required for Production:

1. **Paystack** (Payment Processing)
   ```env
   PAYSTACK_SECRET_KEY=sk_live_your_key
   PAYSTACK_PUBLIC_KEY=pk_live_your_key
   ```

2. **Resend** (Email Notifications)
   ```env
   RESEND_API_KEY=re_your_key
   FROM_EMAIL=noreply@yourdomain.com
   ```

3. **Cloud Storage** (Optional - S3/R2)
   ```env
   STORAGE_TYPE=s3
   AWS_ACCESS_KEY_ID=your_key
   AWS_SECRET_ACCESS_KEY=your_secret
   AWS_BUCKET_NAME=your_bucket
   AWS_REGION=us-east-1
   ```

4. **Security**
   ```env
   JWT_SECRET=generate-a-strong-random-secret
   GIN_MODE=release
   ENVIRONMENT=production
   ```

## Troubleshooting

### libvips not found
```bash
# Install libvips for your OS (see Prerequisites)
```

### Database connection failed
```bash
# Check PostgreSQL is running
sudo systemctl status postgresql

# Verify credentials in .env
```

### Port already in use
```bash
# Change port in .env
PORT=8081
```

### Permission denied on test script
```bash
chmod +x scripts/test_api.sh
```

## Support

- Issues: GitHub Issues
- Email: support@eventtickets.com
- Docs: README.md & API_DOCUMENTATION.md

---

**Happy Coding! 🚀**
