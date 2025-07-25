# HimTalks API Quick Reference

## 🚀 Essential Endpoints

### Public APIs (No Auth Required)
```http
POST /message                    # Submit kritik/saran
POST /songfess                   # Submit songfess
GET  /songfess                   # Get recent songfess (public feed)
GET  /songfess/{id}              # Get specific songfess
GET  /api/spotify/search?q={}    # Search Spotify tracks
GET  /api/spotify/track?id={}    # Get track details
```

### Auth APIs
```http
GET  /auth/google/login          # Start Google OAuth
GET  /auth/google/callback       # OAuth callback
POST /auth/logout                # Logout
GET  /api/protected              # Check auth status
```

### Admin APIs (Auth + Admin Required)
```http
GET  /api/admin/messages         # Get all messages
POST /api/admin/message/delete   # Delete message
GET  /api/admin/songfessAll      # Get all songfess
POST /api/admin/songfess/delete  # Delete songfess
GET  /api/admin/list             # Get admin users
POST /api/admin/addAdmin         # Add admin
POST /api/admin/removeAdmin      # Remove admin
GET  /api/admin/configs          # Get system configs
POST /api/admin/configSongfessDays # Update songfess days
GET  /api/admin/blacklist        # Get blacklisted words
POST /api/admin/blacklist        # Add blacklisted word
POST /api/admin/blacklist/remove # Remove blacklisted word
```

## 📝 Request Examples

### Submit Message
```javascript
fetch('https://api.teknohive.me/message', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    content: 'Your message here',
    sender_name: 'John',
    recipient_name: 'Admin',
    category: 'saran' // or 'kritik'
  })
})
```

### Submit Songfess
```javascript
fetch('https://api.teknohive.me/songfess', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    content: 'Love this song!',
    song_id: '4iV5W9uYEdYUVa79Axb7Rh',
    song_title: 'Shape of You',
    artist: 'Ed Sheeran',
    album_art: 'https://i.scdn.co/image/...',
    preview_url: 'https://p.scdn.co/mp3-preview/...',
    start_time: 30,
    end_time: 60,
    sender_name: 'Alice',
    recipient_name: 'Bob'
  })
})
```

### Admin Delete Message
```javascript
fetch('https://api.teknohive.me/api/admin/message/delete', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  credentials: 'include', // Important for cookie auth
  body: JSON.stringify({ ID: 1 })
})
```

## 🌐 WebSocket Integration
```javascript
const ws = new WebSocket('wss://api.teknohive.me/ws');

ws.onmessage = (event) => {
  const message = JSON.parse(event.data);
  
  switch(message.type) {
    case 'message':
      // New message created
      console.log('New message:', message.data);
      break;
    case 'songfess':
      // New songfess created
      console.log('New songfess:', message.data);
      break;
    case 'delete_message':
      // Message deleted by admin
      console.log('Message deleted:', message.data);
      break;
    case 'delete_songfess':
      // Songfess deleted by admin
      console.log('Songfess deleted:', message.data);
      break;
  }
};
```

## ⚠️ Important Notes

1. **Authentication**: Use `credentials: 'include'` in fetch for admin endpoints
2. **CORS**: Your frontend domain must be in allowed origins
3. **Error Handling**: Always check response status codes
4. **WebSocket**: Handle connection drops and reconnection
5. **Validation**: Check content length limits before submission

## 🔍 Testing Commands

```bash
# Test public message
curl -X POST https://api.teknohive.me/message \
  -H "Content-Type: application/json" \
  -d '{"content":"Test","sender_name":"Test","recipient_name":"Test","category":"saran"}'

# Test auth status (with cookies)
curl -b cookies.txt https://api.teknohive.me/api/protected

# Test admin messages (with cookies)
curl -b cookies.txt https://api.teknohive.me/api/admin/messages
```

## 📄 Related Documentation

- `API_SPECIFICATION.md` - Complete API documentation
- `FRONTEND_TESTING_GUIDE.md` - Detailed testing flows
- `FEATURE_AUDIT_CHECKLIST.md` - Feature checklist
- `ADMIN_API_DOCS.md` - Admin-specific documentation
