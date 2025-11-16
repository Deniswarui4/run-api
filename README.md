# Event Ticketing API

A production-ready event ticketing platform built with Go, featuring RBAC authentication, payment processing, QR code generation, PDF tickets, and comprehensive event management.

## Features

### Core Functionality
- **RBAC Authentication**: Role-based access control with JWT tokens
- **Payment Processing**: Integrated with Paystack for secure payments
- **QR Code Generation**: Automatic QR code generation for tickets
- **PDF Tickets**: Beautiful PDF tickets with QR codes
- **Email Notifications**: Automated emails via Resend
- **Image Processing**: Event image optimization using BIMG
- **Storage**: S3/R2 cloud storage with local fallback
- **Rate Limiting**: Protection against API abuse
- **PostgreSQL Database**: Robust data persistence

### User Roles & Permissions

#### Admin
- Set platform fees and withdrawal processing fees
- Approve/deny revenue withdrawal requests with reasons
- Manage user roles and accounts
- View platform statistics and analytics

#### Moderator
- Review and approve/decline events for publication
- Provide feedback on event submissions
- View moderation statistics

#### Organizer
- Create and manage events
- Upload event images
- Create ticket types with pricing
- Submit events for moderation
- Publish approved events
- View event statistics and revenue
- Request withdrawals

#### Attendee
- Browse and search published events
- Purchase tickets with secure payment
- View ticket history
- Download PDF tickets with QR codes
- Receive email confirmations

## Tech Stack

- **Framework**: Gin Web Framework
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT tokens with bcrypt password hashing
- **Payment**: Paystack
- **Email**: Resend
- **Image Processing**: BIMG (libvips)
- **QR Codes**: go-qrcode
- **PDF Generation**: gofpdf
- **Storage**: AWS S3/Cloudflare R2 with local fallback
- **Rate Limiting**: ulule/limiter

## Project Structure

```
event-ticketing-go-api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── auth/                       # Authentication utilities
│   │   ├── jwt.go                  # JWT token generation/validation
│   │   └── password.go             # Password hashing
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── database/
│   │   └── database.go             # Database initialization & migrations
│   ├── handlers/                   # HTTP request handlers
│   │   ├── admin_handler.go        # Admin endpoints
│   │   ├── attendee_handler.go     # Attendee endpoints
│   │   ├── auth_handler.go         # Authentication endpoints
│   │   ├── moderator_handler.go    # Moderator endpoints
│   │   └── organizer_handler.go    # Organizer endpoints
│   ├── middleware/                 # HTTP middleware
│   │   ├── auth.go                 # Authentication middleware
│   │   ├── cors.go                 # CORS middleware
│   │   ├── error.go                # Error handling
│   │   └── ratelimit.go            # Rate limiting
│   ├── models/                     # Database models
│   │   ├── event.go                # Event & TicketType models
│   │   ├── platform.go             # Platform settings & withdrawals
│   │   ├── ticket.go               # Ticket model
│   │   ├── transaction.go          # Transaction model
│   │   └── user.go                 # User model
│   ├── routes/
│   │   └── routes.go               # Route definitions
│   └── services/                   # Business logic services
│       ├── email.go                # Email service
│       ├── image.go                # Image processing
│       ├── paystack.go             # Payment processing
│       ├── pdf.go                  # PDF generation
│       ├── qrcode.go               # QR code generation
│       └── storage.go              # File storage
├── storage/                        # Local storage (gitignored)
├── .env                            # Environment variables (gitignored)
├── .env.example                    # Environment template
├── .gitignore
├── go.mod
└── README.md
```

## Installation

### Prerequisites

1. **Go 1.21+**
2. **PostgreSQL 12+**
3. **libvips** (for image processing)
   ```bash
   # Ubuntu/Debian
   sudo apt-get install libvips-dev

   # macOS
   brew install vips

   # Fedora/RHEL
   sudo dnf install vips-devel
   ```

### Setup

1. **Clone the repository**
   ```bash
   cd event-ticketing-go-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```bash
   createdb event_ticketing
   ```

4. **Configure environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

5. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```

The API will be available at `http://localhost:8080`

## Configuration

### Required Environment Variables

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=event_ticketing

# JWT
JWT_SECRET=your-super-secret-jwt-key

# Paystack (for payments)
PAYSTACK_SECRET_KEY=sk_test_your_key
PAYSTACK_PUBLIC_KEY=pk_test_your_key

# Resend (for emails)
RESEND_API_KEY=re_your_key
FROM_EMAIL=noreply@yourdomain.com
```

### Optional Environment Variables

```env
# Storage (leave empty for local storage)
STORAGE_TYPE=local
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_BUCKET_NAME=
AWS_REGION=us-east-1
AWS_ENDPOINT=  # For Cloudflare R2

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=1m

# Platform Fees
DEFAULT_PLATFORM_FEE_PERCENTAGE=5.0
DEFAULT_WITHDRAWAL_FEE_PERCENTAGE=2.5
```

## API Documentation

### Authentication

#### Register
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890",
  "role": "attendee"  // or "organizer"
}
```

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "token": "eyJhbGc...",
  "user": { ... }
}
```

### Protected Endpoints

All protected endpoints require the `Authorization` header:
```http
Authorization: Bearer <your-jwt-token>
```

### Admin Endpoints

#### Update Platform Settings
```http
PUT /api/v1/admin/settings
Authorization: Bearer <admin-token>

{
  "platform_fee_percentage": 5.0,
  "withdrawal_fee_percentage": 2.5,
  "min_withdrawal_amount": 1000
}
```

#### Review Withdrawal Request
```http
POST /api/v1/admin/withdrawals/:id/review
Authorization: Bearer <admin-token>

{
  "action": "approve",  // or "reject"
  "comment": "Approved for processing"
}
```

#### Get Platform Statistics
```http
GET /api/v1/admin/stats
Authorization: Bearer <admin-token>
```

### Moderator Endpoints

#### Get Pending Events
```http
GET /api/v1/moderator/events/pending
Authorization: Bearer <moderator-token>
```

#### Review Event
```http
POST /api/v1/moderator/events/:id/review
Authorization: Bearer <moderator-token>

{
  "action": "approve",  // or "reject"
  "comment": "Event meets all guidelines"
}
```

### Organizer Endpoints

#### Create Event
```http
POST /api/v1/organizer/events
Authorization: Bearer <organizer-token>

{
  "title": "Summer Music Festival",
  "description": "Amazing outdoor music festival",
  "category": "Music",
  "venue": "Central Park",
  "address": "123 Park Ave",
  "city": "New York",
  "country": "USA",
  "start_date": "2024-07-15T18:00:00Z",
  "end_date": "2024-07-15T23:00:00Z"
}
```

#### Upload Event Image
```http
POST /api/v1/organizer/events/:id/image
Authorization: Bearer <organizer-token>
Content-Type: multipart/form-data

image: <file>
```

#### Create Ticket Type
```http
POST /api/v1/organizer/events/:id/ticket-types
Authorization: Bearer <organizer-token>

{
  "name": "General Admission",
  "description": "Standard entry ticket",
  "price": 5000,
  "quantity": 500,
  "max_per_order": 10,
  "sale_start": "2024-06-01T00:00:00Z",
  "sale_end": "2024-07-15T18:00:00Z"
}
```

#### Request Withdrawal
```http
POST /api/v1/organizer/withdrawals
Authorization: Bearer <organizer-token>

{
  "amount": 50000,
  "bank_name": "First Bank",
  "account_number": "1234567890",
  "account_name": "John Doe"
}
```

### Attendee Endpoints

#### Browse Events
```http
GET /api/v1/events?category=Music&city=Lagos&search=festival
```

#### Purchase Tickets
```http
POST /api/v1/tickets/purchase
Authorization: Bearer <attendee-token>

{
  "event_id": "uuid",
  "ticket_type_id": "uuid",
  "quantity": 2
}

Response:
{
  "transaction_id": "uuid",
  "payment_reference": "TXN-abc123",
  "authorization_url": "https://checkout.paystack.com/...",
  "amount": 10000,
  "currency": "NGN"
}
```

#### Verify Payment
```http
GET /api/v1/payments/verify?reference=TXN-abc123
```

#### Get My Tickets
```http
GET /api/v1/tickets/my-tickets
Authorization: Bearer <attendee-token>
```

#### Download Ticket PDF
```http
GET /api/v1/tickets/:id/download
Authorization: Bearer <attendee-token>
```

## Database Schema

### Key Tables

- **users**: User accounts with roles
- **events**: Event information and status
- **ticket_types**: Different ticket categories per event
- **tickets**: Individual ticket purchases
- **transactions**: Payment records
- **platform_settings**: Platform configuration
- **withdrawal_requests**: Organizer withdrawal requests
- **organizer_balances**: Organizer earnings tracking

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o event-ticketing-api cmd/api/main.go
```

### Running with Docker
```bash
# Build image
docker build -t event-ticketing-api .

# Run container
docker run -p 8080:8080 --env-file .env event-ticketing-api
```

## Security Features

- **Password Hashing**: bcrypt with salt
- **JWT Authentication**: Secure token-based auth
- **Rate Limiting**: Prevents API abuse
- **CORS**: Configurable cross-origin requests
- **SQL Injection Protection**: GORM parameterized queries
- **Input Validation**: Request validation middleware

## Payment Flow

1. Attendee initiates ticket purchase
2. System creates pending transaction
3. Paystack payment URL generated
4. User completes payment on Paystack
5. System verifies payment with Paystack
6. Tickets created with QR codes and PDFs
7. Email sent to attendee
8. Organizer balance updated (minus platform fee)

## Withdrawal Flow

1. Organizer requests withdrawal
2. Amount deducted from available balance
3. Admin reviews request
4. Admin approves/rejects with comment
5. If approved, admin processes payment
6. Organizer receives email notification

## Event Publishing Flow

1. Organizer creates event (draft status)
2. Organizer adds ticket types
3. Organizer submits for review (pending status)
4. Moderator reviews event
5. Moderator approves/rejects with feedback
6. If approved, organizer can publish (published status)
7. Event appears in public listings

## Storage Options

### Local Storage (Default)
Files stored in `./storage` directory with organized folders:
- `storage/events/` - Event images
- `storage/tickets/qrcodes/` - Ticket QR codes
- `storage/tickets/pdfs/` - Ticket PDFs

### Cloud Storage (S3/R2)
Configure AWS credentials in `.env` to use S3 or Cloudflare R2:
```env
STORAGE_TYPE=s3
AWS_ACCESS_KEY_ID=your_key
AWS_SECRET_ACCESS_KEY=your_secret
AWS_BUCKET_NAME=your_bucket
AWS_REGION=us-east-1
AWS_ENDPOINT=https://your-r2-endpoint  # For R2 only
```

## Troubleshooting

### libvips not found
Install libvips development libraries for your OS (see Prerequisites)

### Database connection failed
Ensure PostgreSQL is running and credentials in `.env` are correct

### Payment verification fails
Check Paystack API keys and ensure callback URL is accessible

### Email not sending
Verify Resend API key and from email address

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - feel free to use this project for personal or commercial purposes.

## Support

For issues and questions:
- Create an issue on GitHub
- Email: support@eventtickets.com

## Roadmap

- [ ] Webhook support for real-time payment notifications
- [ ] Event analytics dashboard
- [ ] Multi-currency support
- [ ] Ticket transfer functionality
- [ ] Mobile app API endpoints
- [ ] Social media integration
- [ ] Discount codes and promotions
- [ ] Recurring events support

---

Built with ❤️ using Go and Gin Framework
