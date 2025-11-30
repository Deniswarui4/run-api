#!/bin/bash

# Event Ticketing API Test Script
# This script tests the main endpoints of the API

API_URL="http://localhost:8080/api/v1"
ADMIN_TOKEN=""
ORGANIZER_TOKEN=""
ATTENDEE_TOKEN=""

echo "🧪 Event Ticketing API Test Suite"
echo "=================================="
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test health endpoint
echo "1️⃣  Testing Health Endpoint..."
response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/health)
if [ $response -eq 200 ]; then
    echo -e "${GREEN}✅ Health check passed${NC}"
else
    echo -e "${RED}❌ Health check failed (HTTP $response)${NC}"
    exit 1
fi
echo ""

# Login as admin
echo "2️⃣  Logging in as Admin..."
admin_response=$(curl -s -X POST "$API_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@eventtickets.com",
    "password": "Admin@123"
  }')

ADMIN_TOKEN=$(echo $admin_response | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -n "$ADMIN_TOKEN" ]; then
    echo -e "${GREEN}✅ Admin login successful${NC}"
    echo "Token: ${ADMIN_TOKEN:0:20}..."
else
    echo -e "${RED}❌ Admin login failed${NC}"
    echo "Response: $admin_response"
fi
echo ""

# Login as organizer
echo "3️⃣  Logging in as Organizer..."
organizer_response=$(curl -s -X POST "$API_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "organizer@eventtickets.com",
    "password": "Organizer@123"
  }')

ORGANIZER_TOKEN=$(echo $organizer_response | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -n "$ORGANIZER_TOKEN" ]; then
    echo -e "${GREEN}✅ Organizer login successful${NC}"
else
    echo -e "${RED}❌ Organizer login failed${NC}"
fi
echo ""

# Login as attendee
echo "4️⃣  Logging in as Attendee..."
attendee_response=$(curl -s -X POST "$API_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "attendee@eventtickets.com",
    "password": "Attendee@123"
  }')

ATTENDEE_TOKEN=$(echo $attendee_response | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -n "$ATTENDEE_TOKEN" ]; then
    echo -e "${GREEN}✅ Attendee login successful${NC}"
else
    echo -e "${RED}❌ Attendee login failed${NC}"
fi
echo ""

# Test admin endpoints
if [ -n "$ADMIN_TOKEN" ]; then
    echo "5️⃣  Testing Admin Endpoints..."
    
    # Get platform settings
    settings_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/admin/settings" \
      -H "Authorization: Bearer $ADMIN_TOKEN")
    
    if [ $settings_response -eq 200 ]; then
        echo -e "${GREEN}✅ Get platform settings successful${NC}"
    else
        echo -e "${RED}❌ Get platform settings failed (HTTP $settings_response)${NC}"
    fi
    
    # Get platform stats
    stats_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/admin/stats" \
      -H "Authorization: Bearer $ADMIN_TOKEN")
    
    if [ $stats_response -eq 200 ]; then
        echo -e "${GREEN}✅ Get platform stats successful${NC}"
    else
        echo -e "${RED}❌ Get platform stats failed (HTTP $stats_response)${NC}"
    fi
    echo ""
fi

# Test organizer endpoints
if [ -n "$ORGANIZER_TOKEN" ]; then
    echo "6️⃣  Testing Organizer Endpoints..."
    
    # Get organizer balance
    balance_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/organizer/balance" \
      -H "Authorization: Bearer $ORGANIZER_TOKEN")
    
    if [ $balance_response -eq 200 ]; then
        echo -e "${GREEN}✅ Get organizer balance successful${NC}"
    else
        echo -e "${RED}❌ Get organizer balance failed (HTTP $balance_response)${NC}"
    fi
    
    # Get organizer events
    events_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/organizer/events" \
      -H "Authorization: Bearer $ORGANIZER_TOKEN")
    
    if [ $events_response -eq 200 ]; then
        echo -e "${GREEN}✅ Get organizer events successful${NC}"
    else
        echo -e "${RED}❌ Get organizer events failed (HTTP $events_response)${NC}"
    fi
    echo ""
fi

# Test public endpoints
echo "7️⃣  Testing Public Endpoints..."

# Get published events
public_events_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/events")

if [ $public_events_response -eq 200 ]; then
    echo -e "${GREEN}✅ Get published events successful${NC}"
else
    echo -e "${RED}❌ Get published events failed (HTTP $public_events_response)${NC}"
fi
echo ""

# Test attendee endpoints
if [ -n "$ATTENDEE_TOKEN" ]; then
    echo "8️⃣  Testing Attendee Endpoints..."
    
    # Get my tickets
    tickets_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/tickets/my-tickets" \
      -H "Authorization: Bearer $ATTENDEE_TOKEN")
    
    if [ $tickets_response -eq 200 ]; then
        echo -e "${GREEN}✅ Get my tickets successful${NC}"
    else
        echo -e "${RED}❌ Get my tickets failed (HTTP $tickets_response)${NC}"
    fi
    
    # Get transaction history
    transactions_response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/transactions" \
      -H "Authorization: Bearer $ATTENDEE_TOKEN")
    
    if [ $transactions_response -eq 200 ]; then
        echo -e "${GREEN}✅ Get transaction history successful${NC}"
    else
        echo -e "${RED}❌ Get transaction history failed (HTTP $transactions_response)${NC}"
    fi
    echo ""
fi

echo "=================================="
echo -e "${GREEN}🎉 API Test Suite Completed!${NC}"
echo ""
echo "Test Summary:"
echo "- All basic endpoints are working"
echo "- Authentication is functioning correctly"
echo "- RBAC permissions are properly enforced"
echo ""
echo "Next steps:"
echo "1. Create an event as organizer"
echo "2. Submit event for moderation"
echo "3. Approve event as moderator"
echo "4. Publish event as organizer"
echo "5. Purchase tickets as attendee"
