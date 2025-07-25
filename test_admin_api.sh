#!/bin/bash

# Admin API Testing Script
# Make sure you're logged in as admin first

API_BASE="https://api.teknohive.me"

echo "=== Admin API Testing Script ==="
echo "Testing admin endpoints with cookie authentication..."
echo

# Test protected endpoint (check if authenticated)
echo "1. Testing authentication..."
curl -s -w "HTTP Status: %{http_code}\n" \
  -H "Content-Type: application/json" \
  --cookie-jar cookies.txt \
  --cookie cookies.txt \
  "$API_BASE/api/protected"
echo

# Test getting all messages
echo "2. Testing GET /api/admin/messages..."
curl -s -w "HTTP Status: %{http_code}\n" \
  -H "Content-Type: application/json" \
  --cookie cookies.txt \
  "$API_BASE/api/admin/messages" | head -20
echo

# Test getting all songfess
echo "3. Testing GET /api/admin/songfessAll..."
curl -s -w "HTTP Status: %{http_code}\n" \
  -H "Content-Type: application/json" \
  --cookie cookies.txt \
  "$API_BASE/api/admin/songfessAll" | head -20
echo

# Test getting songfess with cutoff
echo "4. Testing GET /api/admin/songfess (with cutoff)..."
curl -s -w "HTTP Status: %{http_code}\n" \
  -H "Content-Type: application/json" \
  --cookie cookies.txt \
  "$API_BASE/api/admin/songfess" | head -20
echo

echo "=== Testing completed ==="
echo "Note: To test delete functions, you need to specify a valid ID"
echo "Example delete message: curl -X POST -H 'Content-Type: application/json' --cookie cookies.txt -d '{\"ID\":1}' $API_BASE/api/admin/message/delete"
echo "Example delete songfess: curl -X POST -H 'Content-Type: application/json' --cookie cookies.txt -d '{\"ID\":1}' $API_BASE/api/admin/songfess/delete"
