# HimTalks API Testing Script for PowerShell
# Run with: .\test_api.ps1

$API_BASE = "https://api.teknohive.me"

Write-Host "=== HimTalks API Testing Script ===" -ForegroundColor Green
Write-Host "Testing endpoints..." -ForegroundColor Yellow
Write-Host ""

# Test public message submission
Write-Host "1. Testing public message submission..." -ForegroundColor Cyan
$messageBody = @{
    content = "Test message from PowerShell"
    sender_name = "Tester"
    recipient_name = "Admin"
    category = "saran"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$API_BASE/message" -Method POST -Body $messageBody -ContentType "application/json"
    Write-Host "✅ Message submitted successfully. ID: $($response.id)" -ForegroundColor Green
} catch {
    Write-Host "❌ Message submission failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host ""

# Test songfess feed
Write-Host "2. Testing songfess feed..." -ForegroundColor Cyan
try {
    $songfess = Invoke-RestMethod -Uri "$API_BASE/songfess" -Method GET
    Write-Host "✅ Songfess feed loaded. Count: $($songfess.Count)" -ForegroundColor Green
} catch {
    Write-Host "❌ Songfess feed failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host ""

# Test Spotify search
Write-Host "3. Testing Spotify search..." -ForegroundColor Cyan
try {
    $spotifySearch = Invoke-RestMethod -Uri "$API_BASE/api/spotify/search?q=shape%20of%20you" -Method GET
    Write-Host "✅ Spotify search successful. Results found." -ForegroundColor Green
} catch {
    Write-Host "❌ Spotify search failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host ""

# Test authentication check
Write-Host "4. Testing authentication..." -ForegroundColor Cyan
try {
    $authResponse = Invoke-RestMethod -Uri "$API_BASE/api/protected" -Method GET -WebSession $session
    Write-Host "✅ Authenticated as: $authResponse" -ForegroundColor Green
    $isAuthenticated = $true
} catch {
    Write-Host "⚠️  Not authenticated. Login required for admin features." -ForegroundColor Yellow
    $isAuthenticated = $false
}

Write-Host ""

# Test admin endpoints if authenticated
if ($isAuthenticated) {
    Write-Host "5. Testing admin endpoints..." -ForegroundColor Cyan
    
    # Test admin messages
    try {
        $adminMessages = Invoke-RestMethod -Uri "$API_BASE/api/admin/messages" -Method GET -WebSession $session
        Write-Host "✅ Admin messages loaded. Count: $($adminMessages.Count)" -ForegroundColor Green
    } catch {
        Write-Host "❌ Admin messages failed: $($_.Exception.Message)" -ForegroundColor Red
    }
    
    # Test admin songfess
    try {
        $adminSongfess = Invoke-RestMethod -Uri "$API_BASE/api/admin/songfessAll" -Method GET -WebSession $session
        Write-Host "✅ Admin songfess loaded. Count: $($adminSongfess.Count)" -ForegroundColor Green
    } catch {
        Write-Host "❌ Admin songfess failed: $($_.Exception.Message)" -ForegroundColor Red
    }
} else {
    Write-Host "5. Skipping admin endpoints (not authenticated)" -ForegroundColor Yellow
    Write-Host "   To test admin features:" -ForegroundColor Yellow
    Write-Host "   1. Login via: $API_BASE/auth/google/login" -ForegroundColor Yellow
    Write-Host "   2. Run this script again" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "=== Testing completed ===" -ForegroundColor Green
Write-Host ""
Write-Host "Manual testing URLs:" -ForegroundColor Yellow
Write-Host "- Login: $API_BASE/auth/google/login" -ForegroundColor White
Write-Host "- Logout: $API_BASE/auth/logout (POST)" -ForegroundColor White
Write-Host ""
Write-Host "WebSocket URL: wss://api.teknohive.me/ws" -ForegroundColor White
