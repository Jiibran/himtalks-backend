# Admin API Documentation - HimTalks

## Authentication
All admin endpoints require:
1. **OAuth Google Login**: Login via `/auth/google/login`
2. **Admin Privileges**: User must be registered as admin in database
3. **JWT Cookie**: Authentication via HTTP-only cookie (automatic after login)

**⚠️ Important Setup**: 
- First time setup requires adding admin user to database
- Use: `go run add_admin.go your-email@gmail.com`
- See `ADMIN_TROUBLESHOOTING.md` for detailed setup guide

## Admin Endpoints

### 🔐 Authentication Check
```
GET /api/protected
```
**Description**: Check if user is authenticated
**Response**: User email if authenticated

---

### 📨 Message Management

#### Get All Messages
```
GET /api/admin/messages
```
**Description**: Retrieve all messages (kritik & saran)
**Response**: Array of message objects

**Example Response**:
```json
[
  {
    "id": 1,
    "content": "Great app!",
    "sender_name": "John",
    "recipient_name": "Admin",
    "category": "saran",
    "created_at": "2025-07-25T10:30:00Z"
  }
]
```

#### Delete Message
```
POST /api/admin/message/delete
```
**Body**:
```json
{
  "ID": 1
}
```
**Description**: Delete a message by ID
**Response**: Success/error message

---

### 🎵 Songfess Management

#### Get All Songfess (No Time Limit)
```
GET /api/admin/songfessAll
```
**Description**: Retrieve all songfess regardless of age
**Response**: Array of songfess objects

#### Get Recent Songfess (With Time Limit)
```
GET /api/admin/songfess
```
**Description**: Retrieve songfess within configured time limit (default 7 days)
**Response**: Array of songfess objects

**Example Response**:
```json
[
  {
    "id": 1,
    "content": "Love this song!",
    "song_id": "4iV5W9uYEdYUVa79Axb7Rh",
    "song_title": "Shape of You",
    "artist": "Ed Sheeran",
    "album_art": "https://...",
    "preview_url": "https://...",
    "start_time": 30,
    "end_time": 60,
    "sender_name": "Alice",
    "recipient_name": "Bob",
    "created_at": "2025-07-25T10:30:00Z"
  }
]
```

#### Delete Songfess
```
POST /api/admin/songfess/delete
```
**Body**:
```json
{
  "ID": 1
}
```
**Description**: Delete a songfess by ID
**Response**: Success/error message

---

### 👥 Admin Management

#### Get Admin List
```
GET /api/admin/list
```
**Description**: Get list of all admin users

#### Add New Admin
```
POST /api/admin/addAdmin
```
**Body**:
```json
{
  "Email": "newadmin@example.com"
}
```

#### Remove Admin
```
POST /api/admin/removeAdmin
```
**Body**:
```json
{
  "Email": "admin@example.com"
}
```

---

### ⚙️ Configuration Management

#### Get System Configs
```
GET /api/admin/configs
```
**Description**: Get all system configurations

#### Update Songfess Days Limit
```
POST /api/admin/configSongfessDays
```
**Body**:
```json
{
  "Days": "7"
}
```

---

### 🚫 Blacklist Management

#### Get Blacklisted Words
```
GET /api/admin/blacklist
```

#### Add Blacklisted Word
```
POST /api/admin/blacklist
```
**Body**:
```json
{
  "Word": "badword"
}
```

#### Remove Blacklisted Word
```
POST /api/admin/blacklist/remove
```
**Body**:
```json
{
  "Word": "badword"
}
```

---

## Error Responses

- **401 Unauthorized**: Not logged in or invalid JWT
- **403 Forbidden**: Not an admin user
- **404 Not Found**: Resource not found
- **500 Internal Server Error**: Server error

## Frontend Integration

This backend API is designed to work with a separate frontend application. 
See `API_SPECIFICATION.md` for complete API documentation and `FRONTEND_TESTING_GUIDE.md` for testing flows.

## Testing

Use the provided test script:
```bash
bash test_admin_api.sh
```

## WebSocket Updates

Admin actions trigger real-time updates via WebSocket:
- `delete_message`: When a message is deleted
- `delete_songfess`: When a songfess is deleted
- `message`: When a new message is created
- `songfess`: When a new songfess is created
