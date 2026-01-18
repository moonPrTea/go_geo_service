# go_geo_service

REST API service for managing geo-based incidents and checking whether a user is located in a dangerous zone.
The service supports asynchronous webhook notifications using Redis queue and background worker.

## Features

- CRUD operations for incidents
- Check user location against active danger zones
- API key authentication middleware
- Request statistics endpoint
- Asynchronous webhook delivery via Redis
- Webhook processing
- Postman collection included

---

## Tech Stack

- Go
- Gin
- PostgreSQL
- Redis
- Postman

# Run the service

```bash
go run cmd/api/main.go
```

The API will be available at:
http://localhost:8080/api/v1


## Authentication

Protected endpoints require API key authentication. Add the following header to each request:

```http
X-API-Key: secret-key
```

## API Endpoints

| Method | Endpoint | Requires Auth | Description |
|--------|----------|---------------|-------------|
| GET | `/api/v1/system/health` | Yes | Check system health and availability |
| GET | `/api/v1/incidents` | Yes | Get all incidents |
| GET | `/api/v1/incidents/:id` | Yes | Get incident by ID |
| POST | `/api/v1/incidents` | Yes | Create new incident |
| PUT | `/api/v1/incidents/:id` | Yes | Update incident by ID |
| DELETE | `/api/v1/incidents/:id` | Yes | Delete incident by ID |
| POST | `/api/v1/location/check` | No | Check if user is in danger zone |
| POST | `/api/v1/incidents/stats` | Yes | Get request statistics |h

## Example requests

### 1. All incidents
```bash
curl -X GET \
  http://localhost:8080/api/v1/incidents/ \
  -H "X-API-Key: key"
```

Result:
```json
{
    "incidents": [
        {
            "id": 1,
            "title": "new incident.Incredible",
            "lat": 14.2,
            "lng": 13.1,
            "radius": 1,
            "active": true,
            "created_at": "2026-01-11T15:34:18.251726Z",
            "updated_at": "2026-01-11T15:34:18.251726Z"
        },
        {
            "id": 2,
            "title": "work, holidays r over",
            "lat": 14.2,
            "lng": 13.1,
            "radius": 1,
            "active": true,
            "created_at": "2026-01-17T21:09:49.169952Z",
            "updated_at": "2026-01-18T17:09:03.543817Z"
        }
    ],
    "total": 2
}
```

### 2. Create incident
```bash
curl -X POST \
  http://localhost:8080/api/v1/incidents/ \
  -H "Content-Type: application/json" \
  -H "X-API-Key: key" \
  -d '{
    "title": "new incident.Incredible",
    "lat": 14.2,
    "lng": 13.1,
    "radius": 1,
    "active": true
}'
```

Result:
```json 
{
    "id": 9,
    "title": "new incident.Incredible",
    "lat": 14.2,
    "lng": 13.1,
    "radius": 1,
    "active": true,
    "created_at": "2026-01-18T17:14:08.644565Z",
    "updated_at": "2026-01-18T17:14:08.644565Z"
}
```

### 3. Update incident
```bash
curl -X PUT \
  http://localhost:8080/api/v1/incidents/9 \
  -H "Content-Type: application/json" \
  -H "X-API-Key: key" \
  -d '{
    "title": "work, holidays r over",
    "lat": 14.2,
    "lng": 13.1,
    "radius": 1,
    "active": true
}'
```

Result:
```json 
{
    "message": "Incident data have successfully updated"
}
```

### 4. Get incident by ID
```bash
curl -X GET \
  http://localhost:8080/api/v1/incidents/2 \
  -H "X-API-Key: key"
```

Result:
```json 
{
    "id": 2,
    "title": "work, holidays r over",
    "lat": 14.2,
    "lng": 13.1,
    "radius": 1,
    "active": true,
    "created_at": "2026-01-17T21:09:49.169952Z",
    "updated_at": "2026-01-18T17:09:03.543817Z"
}
```

### 5. Delete incident by ID
```bash
curl -X DELETE \
  http://localhost:8080/api/v1/incidents/2 \
  -H "X-API-Key: key"
```

Result:
```json 
{
    "message": "Incident have successfully deleted"
}
```

### 6. Check location
```bash
curl -X POST \
  http://localhost:8080/api/v1/location/check \
  -H "Content-Type: application/json" \
  -d '{
    "lat": 13.1,
    "lng": 14.2,
    "user_id": "hY*tg3u219hsa"
  }'
```

Result:
```json 
{
    "user_id": "hY*tg3u219hsa",
    "title": "",
    "lat": 13.1,
    "lng": 14.2,
    "zones": null,
    "is_danger": false,
    "timestamp": "2026-01-18T17:17:56.552234+03:00"
}
```

### 7. Get stats
```bash
curl -X POST \
  http://localhost:8080/api/v1/incidents/stats \
  -H "X-API-Key: key"
```

Result:
```json 
{
    "user_count": 1,
    "window_minutes": 60,
    "timestamp": "2026-01-18T17:18:35+03:00"
}
```

### 8. Health request
```bash
curl -X GET \
  http://localhost:8080/api/v1/system/health \
  -H "X-API-Key: key"
```

Result:
```json
{
    "checked_at": 1768744913,
    "status": "ok",
    "system": "active"
}
```

## Webhooks

When a user is detected inside an active incident zone, a webhook payload is sent asynchronously.

### Flow
1. Payload is pushed to Redis queue
2. Background listener consumes messages
3. listener sends HTTP POST request to `WEBHOOK_URL`

### Example Payload

```json
{
  "user_id": "example",
  "latitude": 12.12,
  "longitude": 43.17,
  "zones": ["Cool Place"],
  "timestamp": "2026-01-18T16:04:05Z"
}
```

## HTTP Status Codes

| Code | Description |
|------|-------------|
| `200 OK` | Request successful |
| `201 Created` | Resource created successfully |
| `400 Bad Request` | Invalid request parameters |
| `401 Unauthorized` | Missing or invalid API key |
| `404 Not Found` | Resource not found |
| `500 Internal Server Error` | Internal server error |

## Quick Start

1. **Authentication**: Include `X-API-Key: secret-key` header for protected endpoints
2. **Check Location**: Use `POST /api/v1/location/check` to verify user locations
3. **Manage Incidents**: Create, read, update, and delete incident zones
4. **Receive Webhooks**: Set up a webhook server to receive real-time alerts