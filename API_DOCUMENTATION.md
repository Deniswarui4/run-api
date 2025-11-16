# Event Ticketing API - Complete Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
All protected endpoints require a Bearer token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## 📋 Table of Contents
1. [Authentication Endpoints](#authentication-endpoints)
2. [Admin Endpoints](#admin-endpoints)
3. [Moderator Endpoints](#moderator-endpoints)
4. [Organizer Endpoints](#organizer-endpoints)
5. [Attendee Endpoints](#attendee-endpoints)
6. [Public Endpoints](#public-endpoints)

---

## Authentication Endpoints

### Register User
**POST** `/auth/register`

Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890",
  "role": "attendee"
}
```

**Role Options:** `attendee`, `organizer`

**Response (201):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890",
    "role": "attendee",
    "is_active": true,
    "is_verified": false,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### Login
**POST** `/auth/login`

Authenticate and receive a JWT token.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123"
}
```

**Response (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "attendee"
  }
}
```

### Get Profile
**GET** `/profile`

Get current user's profile.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890",
  "role": "attendee",
  "is_active": true
}
```

### Update Profile
**PUT** `/profile`

Update current user's profile.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "first_name": "Jane",
  "last_name": "Smith",
  "phone": "+0987654321"
}
```

---

## Admin Endpoints

All admin endpoints require `admin` role.

### Get Platform Settings
**GET** `/admin/settings`

Retrieve current platform settings.

**Response (200):**
```json
{
  "id": "uuid",
  "platform_fee_percentage": 5.0,
  "withdrawal_fee_percentage": 2.5,
  "min_withdrawal_amount": 1000.0,
  "currency": "NGN",
  "updated_by": "uuid",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Update Platform Settings
**PUT** `/admin/settings`

Update platform fee settings.

**Request Body:**
```json
{
  "platform_fee_percentage": 5.0,
  "withdrawal_fee_percentage": 2.5,
  "min_withdrawal_amount": 1000.0
}
```

### Get Withdrawal Requests
**GET** `/admin/withdrawals?status=pending`

Get all withdrawal requests. Optional query parameter: `status` (pending, approved, rejected, processed)

**Response (200):**
```json
[
  {
    "id": "uuid",
    "organizer_id": "uuid",
    "amount": 50000,
    "withdrawal_fee": 1250,
    "net_amount": 48750,
    "status": "pending",
    "bank_name": "First Bank",
    "account_number": "1234567890",
    "account_name": "John Doe",
    "created_at": "2024-01-01T00:00:00Z",
    "organizer": {
      "id": "uuid",
      "email": "organizer@example.com",
      "first_name": "John",
      "last_name": "Doe"
    }
  }
]
```

### Review Withdrawal Request
**POST** `/admin/withdrawals/:id/review`

Approve or reject a withdrawal request.

**Request Body:**
```json
{
  "action": "approve",
  "comment": "Approved for processing"
}
```

**Action Options:** `approve`, `reject`

### Process Withdrawal
**POST** `/admin/withdrawals/:id/process`

Mark an approved withdrawal as processed.

**Request Body:**
```json
{
  "transaction_ref": "TXN-BANK-123456"
}
```

### Get Platform Statistics
**GET** `/admin/stats`

Get platform-wide statistics.

**Response (200):**
```json
{
  "total_users": 1250,
  "total_organizers": 45,
  "total_events": 120,
  "total_tickets_sold": 5430,
  "total_revenue": 2715000.0,
  "platform_revenue": 135750.0
}
```

### Get All Users
**GET** `/admin/users?role=organizer&page=1&limit=20`

Get all users with optional filtering.

**Query Parameters:**
- `role`: Filter by role (admin, moderator, organizer, attendee)
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)

### Manage User Role
**PUT** `/admin/users/:id/role`

Change a user's role.

**Request Body:**
```json
{
  "role": "moderator"
}
```

### Toggle User Status
**PUT** `/admin/users/:id/toggle-status`

Activate or deactivate a user account.

---

## Moderator Endpoints

All moderator endpoints require `moderator` or `admin` role.

### Get Pending Events
**GET** `/moderator/events/pending`

Get all events pending moderation.

**Response (200):**
```json
[
  {
    "id": "uuid",
    "title": "Summer Music Festival",
    "description": "Amazing outdoor music festival",
    "category": "Music",
    "venue": "Central Park",
    "start_date": "2024-07-15T18:00:00Z",
    "end_date": "2024-07-15T23:00:00Z",
    "status": "pending",
    "organizer": {
      "id": "uuid",
      "email": "organizer@example.com",
      "first_name": "John",
      "last_name": "Doe"
    }
  }
]
```

### Get Event for Review
**GET** `/moderator/events/:id`

Get detailed information about an event for review.

### Review Event
**POST** `/moderator/events/:id/review`

Approve or reject an event.

**Request Body:**
```json
{
  "action": "approve",
  "comment": "Event meets all guidelines"
}
```

**Action Options:** `approve`, `reject`

### Get Moderation Stats
**GET** `/moderator/stats`

Get moderation statistics.

**Response (200):**
```json
{
  "pending_events": 15,
  "approved_events": 120,
  "rejected_events": 8,
  "my_reviews": 45
}
```

### Get My Reviews
**GET** `/moderator/reviews`

Get events reviewed by the current moderator.

---

## Organizer Endpoints

All organizer endpoints require `organizer`, `moderator`, or `admin` role.

### Create Event
**POST** `/organizer/events`

Create a new event.

**Request Body:**
```json
{
  "title": "Summer Music Festival",
  "description": "Amazing outdoor music festival with top artists",
  "category": "Music",
  "venue": "Central Park",
  "address": "123 Park Avenue",
  "city": "New York",
  "country": "USA",
  "start_date": "2024-07-15T18:00:00Z",
  "end_date": "2024-07-15T23:00:00Z"
}
```

**Response (201):**
```json
{
  "id": "uuid",
  "title": "Summer Music Festival",
  "status": "draft",
  "organizer_id": "uuid",
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Upload Event Image
**POST** `/organizer/events/:id/image`

Upload an image for an event.

**Content-Type:** `multipart/form-data`

**Form Data:**
- `image`: Image file (JPEG, PNG, GIF)

**Response (200):**
```json
{
  "image_url": "/storage/events/event-abc123-20240101-150405-xyz789.jpg"
}
```

### Update Event
**PUT** `/organizer/events/:id`

Update an event (only draft or rejected events).

**Request Body:** Same as Create Event

### Get My Events
**GET** `/organizer/events?status=published`

Get organizer's events with optional status filter.

**Query Parameters:**
- `status`: Filter by status (draft, pending, approved, rejected, published, cancelled)

### Submit Event for Review
**POST** `/organizer/events/:id/submit`

Submit an event for moderation.

**Response (200):**
```json
{
  "message": "Event submitted for review",
  "event": {
    "id": "uuid",
    "title": "Summer Music Festival",
    "status": "pending"
  }
}
```

### Publish Event
**POST** `/organizer/events/:id/publish`

Publish an approved event.

### Create Ticket Type
**POST** `/organizer/events/:id/ticket-types`

Create a ticket type for an event.

**Request Body:**
```json
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

### Get Organizer Balance
**GET** `/organizer/balance`

Get organizer's earnings and balance.

**Response (200):**
```json
{
  "id": "uuid",
  "organizer_id": "uuid",
  "total_earnings": 250000,
  "available_balance": 200000,
  "pending_balance": 30000,
  "withdrawn_amount": 20000
}
```

### Request Withdrawal
**POST** `/organizer/withdrawals`

Request a withdrawal of earnings.

**Request Body:**
```json
{
  "amount": 50000,
  "bank_name": "First Bank",
  "account_number": "1234567890",
  "account_name": "John Doe"
}
```

### Get My Withdrawals
**GET** `/organizer/withdrawals`

Get organizer's withdrawal history.

### Get Event Stats
**GET** `/organizer/events/:id/stats`

Get statistics for a specific event.

**Response (200):**
```json
{
  "total_tickets_sold": 450,
  "total_revenue": 225000,
  "net_revenue": 213750,
  "checked_in_tickets": 380
}
```

---

## Attendee Endpoints

### Purchase Tickets
**POST** `/tickets/purchase`

Initiate a ticket purchase.

**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "event_id": "uuid",
  "ticket_type_id": "uuid",
  "quantity": 2
}
```

**Response (200):**
```json
{
  "transaction_id": "uuid",
  "payment_reference": "TXN-abc12345",
  "authorization_url": "https://checkout.paystack.com/...",
  "amount": 10000,
  "currency": "NGN"
}
```

### Verify Payment
**GET** `/payments/verify?reference=TXN-abc12345`

Verify a payment after Paystack redirect.

**Response (200):**
```json
{
  "message": "Payment verified successfully",
  "status": "success",
  "tickets": [
    {
      "id": "uuid",
      "ticket_number": "TKT-abc12345",
      "event_id": "uuid",
      "status": "confirmed",
      "price": 5000,
      "qr_code_url": "/storage/tickets/qrcodes/qr-...",
      "pdf_url": "/storage/tickets/pdfs/ticket-..."
    }
  ]
}
```

### Get My Tickets
**GET** `/tickets/my-tickets`

Get all tickets purchased by the attendee.

**Headers:** `Authorization: Bearer <token>`

**Response (200):**
```json
[
  {
    "id": "uuid",
    "ticket_number": "TKT-abc12345",
    "status": "confirmed",
    "price": 5000,
    "qr_code_url": "/storage/tickets/qrcodes/...",
    "pdf_url": "/storage/tickets/pdfs/...",
    "event": {
      "id": "uuid",
      "title": "Summer Music Festival",
      "venue": "Central Park",
      "start_date": "2024-07-15T18:00:00Z"
    },
    "ticket_type": {
      "name": "General Admission"
    }
  }
]
```

### Get Ticket Details
**GET** `/tickets/:id`

Get details of a specific ticket.

### Download Ticket PDF
**GET** `/tickets/:id/download`

Download the ticket PDF.

**Response:** PDF file download

### Get Transaction History
**GET** `/transactions`

Get attendee's transaction history.

---

## Public Endpoints

### Get Published Events
**GET** `/events?category=Music&city=Lagos&search=festival`

Browse all published events.

**Query Parameters:**
- `category`: Filter by category
- `city`: Filter by city
- `search`: Search in title and description

**Response (200):**
```json
[
  {
    "id": "uuid",
    "title": "Summer Music Festival",
    "description": "Amazing outdoor music festival",
    "category": "Music",
    "venue": "Central Park",
    "city": "New York",
    "image_url": "/storage/events/...",
    "start_date": "2024-07-15T18:00:00Z",
    "end_date": "2024-07-15T23:00:00Z",
    "status": "published",
    "organizer": {
      "first_name": "John",
      "last_name": "Doe"
    },
    "ticket_types": [
      {
        "id": "uuid",
        "name": "General Admission",
        "price": 5000,
        "quantity": 500,
        "sold": 450
      }
    ]
  }
]
```

### Get Event Details
**GET** `/events/:id`

Get detailed information about a specific published event.

---

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
  "error": "Invalid request data"
}
```

### 401 Unauthorized
```json
{
  "error": "Authorization header required"
}
```

### 403 Forbidden
```json
{
  "error": "Insufficient permissions"
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 429 Too Many Requests
```json
{
  "error": "Rate limit exceeded",
  "retry_after": 30
}
```

### 500 Internal Server Error
```json
{
  "error": "An internal error occurred"
}
```

---

## Rate Limiting

The API implements rate limiting to prevent abuse:
- Default: 100 requests per minute per IP
- Rate limit headers are included in responses:
  - `X-RateLimit-Limit`: Maximum requests allowed
  - `X-RateLimit-Remaining`: Remaining requests
  - `X-RateLimit-Reset`: Time when limit resets

---

## Webhooks (Future Implementation)

Webhook endpoints for real-time notifications:
- Payment confirmations
- Event approvals
- Ticket purchases
- Withdrawal status updates

---

## Testing

Use the provided test script:
```bash
./scripts/test_api.sh
```

Or use tools like:
- Postman
- Insomnia
- cURL
- HTTPie

---

## Support

For issues or questions:
- GitHub Issues
- Email: support@eventtickets.com
- Documentation: README.md
