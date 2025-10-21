# Schools CRUD API Documentation

## School Management Endpoints

### Base URL
```
/api/v1/schools
```

## Public Endpoints (No Authentication Required)

### 1. Get All Schools
Get list of all active schools.

#### URL
```
GET /api/v1/schools
```

#### Query Parameters
- `search` (optional): Search schools by name

#### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "มหาวิทยาลัยกรุงเทพ",
      "address": "119 ถนนรังสิต เขตทุ่งสองห้อง กรุงเทพฯ",
      "phone": "02-350-3500",
      "email": "info@bu.ac.th",
      "website_url": "https://www.bu.ac.th",
      "logo_url": null,
      "is_active": true,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-15T14:30:00Z"
    }
  ]
}
```

### 2. Search Schools
Search schools by name.

#### URL
```
GET /api/v1/schools?search={query}
```

#### Example
```
GET /api/v1/schools?search=กรุงเทพ
```

### 3. Get School by ID
Get specific school information.

#### URL
```
GET /api/v1/schools/{id}
```

#### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "มหาวิทยาลัยกรุงเทพ",
    "address": "119 ถนนรังสิต เขตทุ่งสองห้อง กรุงเทพฯ",
    "phone": "02-350-3500",
    "email": "info@bu.ac.th",
    "website_url": "https://www.bu.ac.th",
    "logo_url": null,
    "is_active": true,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-15T14:30:00Z"
  }
}
```

## Protected Endpoints (Authentication Required)

### 4. Create School
Create a new school (requires authentication).

#### URL
```
POST /api/v1/schools
```

#### Headers
```
Authorization: Bearer {access_token}
Content-Type: application/json
```

#### Request Body
```json
{
  "name": "มหาวิทยาลัยเทคโนโลยีราชมงคลกรุงเทพ",
  "address": "2 ถนนนางลิ้นจี่ เขตทุ่งมหาเมฆ กรุงเทพฯ 10140",
  "phone": "02-287-9600",
  "email": "info@rmutk.ac.th",
  "website_url": "https://www.rmutk.ac.th"
}
```

#### Response Format (Success - 201 Created)
```json
{
  "success": true,
  "data": {
    "id": "456e7890-e89b-12d3-a456-426614174001",
    "name": "มหาวิทยาลัยเทคโนโลยีราชมงคลกรุงเทพ",
    "address": "2 ถนนนางลิ้นจี่ เขตทุ่งมหาเมฆ กรุงเทพฯ 10140",
    "phone": "02-287-9600",
    "email": "info@rmutk.ac.th",
    "website_url": "https://www.rmutk.ac.th",
    "logo_url": null,
    "is_active": true,
    "created_at": "2024-01-20T10:00:00Z",
    "updated_at": "2024-01-20T10:00:00Z"
  }
}
```

### 5. Update School
Update school information (requires authentication).

#### URL
```
PATCH /api/v1/schools/{id}
```

#### Headers
```
Authorization: Bearer {access_token}
Content-Type: application/json
```

#### Request Body
```json
{
  "name": "มหาวิทยาลัยกรุงเทพ (สาขาใหม่)",
  "phone": "02-350-3501",
  "website_url": "https://www.bu.ac.th/new-branch"
}
```

#### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "มหาวิทยาลัยกรุงเทพ (สาขาใหม่)",
    "address": "119 ถนนรังสิต เขตทุ่งสองห้อง กรุงเทพฯ",
    "phone": "02-350-3501",
    "email": "info@bu.ac.th",
    "website_url": "https://www.bu.ac.th/new-branch",
    "logo_url": null,
    "is_active": true,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-20T15:30:00Z"
  }
}
```

### 6. Delete School
Soft delete a school (requires authentication).

#### URL
```
DELETE /api/v1/schools/{id}
```

#### Headers
```
Authorization: Bearer {access_token}
```

#### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "message": "School deleted successfully"
  }
}
```

## System Information Endpoint

### Get System Info
Get comprehensive system information and API documentation.

#### URL
```
GET /api/v1/info
```

#### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "service_name": "Easy Attend Service",
    "version": "1.0.0",
    "description": "A comprehensive attendance management system for educational institutions",
    "api_version": "v1",
    "documentation": "/docs",
    "health_check": "/health",
    "features": [
      "User Authentication & Authorization",
      "School Management",
      "Student Registration",
      "Teacher Management",
      "Attendance Tracking",
      "Profile Management",
      "Gender & Prefix Support",
      "Multi-language Support (Thai/English)"
    ],
    "endpoints": {
      "auth": {
        "login": "POST /api/v1/auth/login",
        "register": "POST /api/v1/auth/register",
        "refresh": "POST /api/v1/auth/refresh"
      },
      "profile": {
        "get_own": "GET /api/v1/profile",
        "update": "PATCH /api/v1/profile",
        "get_by_id": "GET /api/v1/profile/{id}"
      },
      "schools": {
        "list": "GET /api/v1/schools",
        "get": "GET /api/v1/schools/{id}",
        "create": "POST /api/v1/schools",
        "update": "PATCH /api/v1/schools/{id}",
        "delete": "DELETE /api/v1/schools/{id}",
        "search": "GET /api/v1/schools?search={query}"
      },
      "reference": {
        "genders": "GET /api/v1/genders",
        "prefixes": "GET /api/v1/prefixes"
      }
    },
    "contact": {
      "developer": "Easy Attend Development Team",
      "email": "support@easyattend.com"
    }
  }
}
```

## Example Usage

### PowerShell Examples
```powershell
# Get all schools
$schools = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/schools" -Method GET
$schools | ConvertTo-Json -Depth 5

# Search schools
$searchResults = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/schools?search=กรุงเทพ" -Method GET
$searchResults | ConvertTo-Json -Depth 5

# Get system info
$info = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/info" -Method GET
$info | ConvertTo-Json -Depth 10

# Create school (requires token)
$schoolData = @{
    name = "มหาวิทยาลัยเทคโนโลยีกิง มองกุฎ"
    address = "126 ถนนประชาราษฎร์ 1 แขวงทุ่งวัดดอน เขตสาทร กรุงเทพฯ"
    phone = "02-287-9999"
    email = "info@kmitl.ac.th"
    website_url = "https://www.kmitl.ac.th"
} | ConvertTo-Json

$headers = @{ 
    Authorization = "Bearer YOUR_ACCESS_TOKEN"
    'Content-Type' = 'application/json'
}

$response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/schools" -Method POST -Body $schoolData -Headers $headers
$response | ConvertTo-Json -Depth 5
```

### cURL Examples
```bash
# Get all schools
curl -X GET "http://localhost:8080/api/v1/schools"

# Search schools
curl -X GET "http://localhost:8080/api/v1/schools?search=กรุงเทพ"

# Get system info
curl -X GET "http://localhost:8080/api/v1/info"

# Create school
curl -X POST "http://localhost:8080/api/v1/schools" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "มหาวิทยาลัยเทคโนโลยีกิง มองกุฎ",
    "address": "126 ถนนประชาราษฎร์ 1 แขวงทุ่งวัดดอน เขตสาทร กรุงเทพฯ",
    "phone": "02-287-9999",
    "email": "info@kmitl.ac.th",
    "website_url": "https://www.kmitl.ac.th"
  }'

# Update school
curl -X PATCH "http://localhost:8080/api/v1/schools/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "02-287-9998",
    "website_url": "https://www.kmitl.ac.th/updated"
  }'

# Delete school
curl -X DELETE "http://localhost:8080/api/v1/schools/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## Error Responses

### School Name Already Exists (409 Conflict)
```json
{
  "code": 409,
  "message": "School name already exists"
}
```

### School Not Found (404 Not Found)
```json
{
  "code": 404,
  "message": "School not found"
}
```

### Cannot Delete School with Users (409 Conflict)
```json
{
  "code": 409,
  "message": "Cannot delete school that has active users"
}
```

### Unauthorized (401 Unauthorized)
```json
{
  "code": 401,
  "message": "User not authenticated"
}
```

## Business Rules

1. **School Name Uniqueness**: School names must be unique (case-insensitive)
2. **Soft Delete**: Schools are soft deleted (is_active = false) rather than permanently deleted
3. **User Protection**: Cannot delete schools that have active users
4. **Public Access**: Anyone can view schools and search them
5. **Protected Operations**: Create, Update, and Delete require authentication
6. **Search Limit**: Search results are limited to 20 schools to prevent large responses

## Features

- ✅ Full CRUD operations
- ✅ Search functionality
- ✅ Soft delete protection
- ✅ Data validation
- ✅ Public read access
- ✅ Protected write operations
- ✅ System information endpoint
- ✅ Comprehensive error handling