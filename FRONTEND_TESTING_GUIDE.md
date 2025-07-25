# HimTalks Frontend Testing Flow & Prompts

## 🧪 Testing Flow for Frontend Integration

### Phase 1: Public User Features Testing

#### 1.1 Anonymous Message (Kritik/Saran)
**Test Flow:**
```
1. User opens message form
2. User fills: content, sender_name, recipient_name, category
3. User submits form
4. Check: API POST /message
5. Verify: Message appears in real-time (WebSocket)
6. Verify: Message content is filtered for blacklisted words
```

**API Test Command:**
```bash
curl -X POST https://api.teknohive.me/message \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Test message",
    "sender_name": "John",
    "recipient_name": "Admin",
    "category": "saran"
  }'
```

#### 1.2 Songfess Submission
**Test Flow:**
```
1. User searches for song (Spotify integration)
2. User selects song from search results
3. User writes message content
4. User fills sender/recipient names
5. User submits songfess
6. Check: API POST /songfess
7. Verify: Songfess appears in public feed
8. Verify: Real-time update via WebSocket
```

**API Test Commands:**
```bash
# Search for songs
curl "https://api.teknohive.me/api/spotify/search?q=shape%20of%20you"

# Get track details
curl "https://api.teknohive.me/api/spotify/track?id=4iV5W9uYEdYUVa79Axb7Rh"

# Submit songfess
curl -X POST https://api.teknohive.me/songfess \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Love this song!",
    "song_id": "4iV5W9uYEdYUVa79Axb7Rh",
    "song_title": "Shape of You",
    "artist": "Ed Sheeran",
    "album_art": "https://i.scdn.co/image/...",
    "preview_url": "https://p.scdn.co/mp3-preview/...",
    "start_time": 30,
    "end_time": 60,
    "sender_name": "Alice",
    "recipient_name": "Bob"
  }'
```

#### 1.3 Public Songfess Feed
**Test Flow:**
```
1. User visits songfess feed page
2. Check: API GET /songfess (recent, within time limit)
3. Verify: Only shows songfess within configured days
4. User clicks on specific songfess
5. Check: API GET /songfess/{id}
6. Verify: Shows full songfess details
7. Verify: Music preview plays correctly
```

**API Test Commands:**
```bash
# Get recent songfess feed
curl https://api.teknohive.me/songfess

# Get specific songfess
curl https://api.teknohive.me/songfess/1
```

### Phase 2: Admin Authentication Testing

#### 2.1 Google OAuth Login
**Test Flow:**
```
1. Admin clicks "Login with Google"
2. Check: Redirects to /auth/google/login
3. Verify: Google OAuth consent screen appears
4. Admin grants permission
5. Check: Callback to /auth/google/callback
6. Verify: JWT cookie is set
7. Check: Redirect to frontend admin panel
```

#### 2.2 Admin Authentication Check
**Test Flow:**
```
1. Admin accesses admin features
2. Check: API GET /api/protected
3. Verify: Returns admin email
4. If failed: Redirect to login
```

**API Test Command:**
```bash
# Check authentication (with cookie)
curl -b cookies.txt https://api.teknohive.me/api/protected
```

### Phase 3: Admin Panel Features Testing

#### 3.1 View All Messages
**Test Flow:**
```
1. Admin opens messages management page
2. Check: API GET /api/admin/messages
3. Verify: Shows all messages (kritik & saran)
4. Verify: Displays message details (content, category, names, timestamp)
5. Verify: Shows total count
```

#### 3.2 Delete Messages
**Test Flow:**
```
1. Admin clicks delete on a message
2. Show confirmation dialog
3. If confirmed: API POST /api/admin/message/delete
4. Verify: Message is removed from list
5. Verify: Success notification shown
6. Verify: Real-time update (WebSocket)
```

#### 3.3 View All Songfess
**Test Flow:**
```
1. Admin opens songfess management page
2. Check: API GET /api/admin/songfessAll (all, no time limit)
3. Verify: Shows all songfess regardless of age
4. Verify: Shows songfess details (content, song info, album art)
5. Verify: Shows total count
```

#### 3.4 Delete Songfess
**Test Flow:**
```
1. Admin clicks delete on a songfess
2. Show confirmation dialog  
3. If confirmed: API POST /api/admin/songfess/delete
4. Verify: Songfess is removed from list
5. Verify: Success notification shown
6. Verify: Real-time update (WebSocket)
```

#### 3.5 Admin User Management
**Test Flow:**
```
1. Admin opens admin management page
2. Check: API GET /api/admin/list
3. Verify: Shows list of admin users
4. Admin adds new admin: API POST /api/admin/addAdmin
5. Admin removes admin: API POST /api/admin/removeAdmin
6. Verify: Cannot remove self
```

#### 3.6 System Configuration
**Test Flow:**
```
1. Admin opens settings page
2. Check: API GET /api/admin/configs  
3. Verify: Shows current configurations
4. Admin updates songfess days: API POST /api/admin/configSongfessDays
5. Verify: Configuration is updated
```

#### 3.7 Blacklist Management
**Test Flow:**
```
1. Admin opens blacklist management
2. Check: API GET /api/admin/blacklist
3. Verify: Shows blacklisted words
4. Admin adds word: API POST /api/admin/blacklist
5. Admin removes word: API POST /api/admin/blacklist/remove
6. Test: Submit message with blacklisted word (should be rejected)
```

### Phase 4: Real-time Features Testing

#### 4.1 WebSocket Connection
**Test Flow:**
```
1. Frontend connects to wss://api.teknohive.me/ws
2. Verify: Connection established successfully
3. Submit new message/songfess
4. Verify: Real-time update received via WebSocket
5. Admin deletes item
6. Verify: Delete notification received
```

**WebSocket Message Types to Handle:**
- `message`: New message created
- `songfess`: New songfess created
- `delete_message`: Message deleted
- `delete_songfess`: Songfess deleted

---

## 🔍 Feature Checklist for Frontend

### ✅ Public Features
- [ ] **Message Form**: Submit kritik/saran with validation
- [ ] **Spotify Search**: Search and select songs
- [ ] **Songfess Form**: Submit songfess with song details
- [ ] **Public Feed**: Display recent songfess
- [ ] **Songfess Details**: View individual songfess
- [ ] **Music Preview**: Play song previews
- [ ] **Real-time Updates**: WebSocket integration for live updates

### ✅ Admin Features  
- [ ] **Google OAuth Login**: Login with Google
- [ ] **Authentication Check**: Verify admin status
- [ ] **Message Management**: View and delete all messages
- [ ] **Songfess Management**: View and delete all songfess
- [ ] **Admin User Management**: Add/remove admin users
- [ ] **System Configuration**: Update settings (songfess days limit)
- [ ] **Blacklist Management**: Manage blacklisted words
- [ ] **Real-time Admin Updates**: Live updates for admin actions
- [ ] **Logout**: Clear authentication

### ✅ UI/UX Features
- [ ] **Error Handling**: Display API errors appropriately
- [ ] **Loading States**: Show loading indicators
- [ ] **Success Notifications**: Confirm successful actions
- [ ] **Confirmation Dialogs**: Confirm destructive actions
- [ ] **Responsive Design**: Mobile-friendly interface
- [ ] **Form Validation**: Client-side validation
- [ ] **Authentication State**: Handle login/logout states

---

## 🚨 Common Issues to Check

1. **CORS Issues**: Ensure your frontend domain is in allowed origins
2. **Cookie Handling**: Ensure credentials are included in requests
3. **Authentication Flow**: Handle redirect after OAuth
4. **WebSocket Connection**: Handle connection drops and reconnection
5. **Error Boundaries**: Handle API errors gracefully
6. **Character Limits**: Respect message content limits
7. **Blacklist Filtering**: Handle blacklisted word errors
8. **Admin Privileges**: Check admin status before showing admin features

---

## 📝 Testing Script Template

```javascript
// Frontend API Testing Template
const API_BASE = 'https://api.teknohive.me';

// Test authentication
async function testAuth() {
  const response = await fetch(`${API_BASE}/api/protected`, {
    credentials: 'include'
  });
  console.log('Auth status:', response.status);
}

// Test message submission
async function testMessage() {
  const response = await fetch(`${API_BASE}/message`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      content: 'Test message',
      sender_name: 'Tester',
      recipient_name: 'Admin',
      category: 'saran'
    })
  });
  console.log('Message submit:', response.status);
}

// Test admin features
async function testAdminMessages() {
  const response = await fetch(`${API_BASE}/api/admin/messages`, {
    credentials: 'include'
  });
  const data = await response.json();
  console.log('Admin messages:', data.length);
}
```
