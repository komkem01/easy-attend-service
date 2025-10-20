# Updated Login and Profile API Responses

## Login Response
Enhanced login response with prefix and gender information.

### URL
```
POST /api/v1/auth/login
```

### Request
```json
{
  "email": "somchai@example.com",
  "password": "password123"
}
```

### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "username": "somchai_jaidee",
      "email": "somchai@example.com",
      "first_name": "สมชาย",
      "last_name": "ใจดี",
      "full_name": "สมชาย ใจดี",
      "role": "student",
      "school_id": "123e4567-e89b-12d3-a456-426614174000",
      "is_active": true,
      "created_at": "2024-01-01T10:00:00Z",
      "prefix": {
        "id": 1,
        "code": "MR",
        "name_th": "นาย",
        "name_en": "Mr.",
        "abbreviation": "นาย"
      },
      "gender": {
        "id": 1,
        "code": "M",
        "name_th": "ชาย",
        "name_en": "Male",
        "abbreviation": "M"
      },
      "school": {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "มหาวิทยาลัยกรุงเทพ"
      }
    }
  }
}
```

## Get Profile Response (Own Profile)
Enhanced profile response with complete user information including prefix and gender.

### URL
```
GET /api/v1/profile
```

### Headers
```
Authorization: Bearer {access_token}
```

### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "somchai_jaidee",
    "email": "somchai@example.com",
    "first_name": "สมชาย",
    "last_name": "ใจดี",
    "full_name": "สมชาย ใจดี",
    "role": "student",
    "is_active": true,
    "created_at": "2024-01-01T10:00:00Z",
    "phone": "0812345678",
    "date_of_birth": "1995-05-15",
    "address": "123 ถนนสุขุมวิท กรุงเทพฯ",
    "prefix": {
      "id": 1,
      "code": "MR",
      "name_th": "นาย",
      "name_en": "Mr.",
      "abbreviation": "นาย"
    },
    "gender": {
      "id": 1,
      "code": "M",
      "name_th": "ชาย",
      "name_en": "Male",
      "abbreviation": "M"
    },
    "school": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "มหาวิทยาลัยกรุงเทพ"
    },
    "profile": {
      "id": "789e1234-e89b-12d3-a456-426614174567",
      "user_id": "550e8400-e29b-41d4-a716-446655440000",
      "full_name": "นายสมชาย ใจดี",
      "bio": "นักศึกษาวิทยาการคอมพิวเตอร์ ปี 3",
      "website": "https://somchai.dev",
      "profile_picture": "https://example.com/profile.jpg",
      "phone_number": "0812345678",
      "date_of_birth": "1995-05-15",
      "address": "123 ถนนสุขุมวิท กรุงเทพฯ",
      "city": "กรุงเทพฯ",
      "state": "กรุงเทพมหานคร",
      "postal_code": "10110",
      "country": "ประเทศไทย",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-15T14:30:00Z"
    }
  }
}
```

## Get Profile by ID Response (Public Profile)
Public profile information with prefix and gender for viewing other users.

### URL
```
GET /api/v1/profile/{user_id}
```

### Response Format (Success - 200 OK)
```json
{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "username": "somchai_jaidee",
    "first_name": "สมชาย",
    "last_name": "ใจดี",
    "full_name": "สมชาย ใจดี",
    "role": "student",
    "created_at": "2024-01-01T10:00:00Z",
    "prefix": {
      "id": 1,
      "code": "MR",
      "name_th": "นาย",
      "name_en": "Mr.",
      "abbreviation": "นาย"
    },
    "gender": {
      "id": 1,
      "code": "M",
      "name_th": "ชาย",
      "name_en": "Male",
      "abbreviation": "M"
    },
    "school": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "มหาวิทยาลัยกรุงเทพ"
    },
    "profile": {
      "id": "789e1234-e89b-12d3-a456-426614174567",
      "user_id": "550e8400-e29b-41d4-a716-446655440000",
      "full_name": "นายสมชาย ใจดี",
      "bio": "นักศึกษาวิทยาการคอมพิวเตอร์",
      "website": "https://somchai.dev",
      "profile_picture": "https://example.com/profile.jpg",
      "city": "กรุงเทพฯ",
      "country": "ประเทศไทย"
    }
  }
}
```

## Key Improvements

### 1. **Login Response Enhancements**
- ✅ Added `prefix` object with complete information
- ✅ Added `gender` object with complete information  
- ✅ Added `created_at` timestamp
- ✅ Maintained backward compatibility

### 2. **Profile Response Enhancements**
- ✅ Added `prefix` and `gender` objects to own profile
- ✅ Added complete user information (`phone`, `date_of_birth`, `address`)
- ✅ Enhanced public profile with prefix/gender for other users
- ✅ Maintained privacy (no sensitive data in public profiles)

### 3. **Data Structure**
Each `prefix` and `gender` object includes:
```json
{
  "id": 1,                    // Database ID
  "code": "MR",              // Short code
  "name_th": "นาย",          // Thai name
  "name_en": "Mr.",          // English name  
  "abbreviation": "นาย"      // Abbreviation
}
```

## Example Usage

### PowerShell Examples
```powershell
# Login
$loginData = @{
    email = "somchai@example.com"
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method POST -Body $loginData -ContentType "application/json"
$token = $loginResponse.data.access_token

# Get own profile
$headers = @{ Authorization = "Bearer $token" }
$profile = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/profile" -Method GET -Headers $headers
$profile | ConvertTo-Json -Depth 10

# Get public profile by ID
$userId = "550e8400-e29b-41d4-a716-446655440000"
$publicProfile = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/profile/$userId" -Method GET
$publicProfile | ConvertTo-Json -Depth 10
```

### cURL Examples
```bash
# Login
curl -X POST "http://localhost:8080/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"somchai@example.com","password":"password123"}'

# Get own profile (replace TOKEN with actual access_token)
curl -X GET "http://localhost:8080/api/v1/profile" \
  -H "Authorization: Bearer TOKEN"

# Get public profile
curl -X GET "http://localhost:8080/api/v1/profile/550e8400-e29b-41d4-a716-446655440000"
```

## Privacy Notes

### Own Profile (`GET /api/v1/profile`)
- ✅ Complete user information
- ✅ Email, phone, address included
- ✅ Complete profile data

### Public Profile (`GET /api/v1/profile/{id}`)
- ❌ No email, phone, or address
- ✅ Basic information only
- ✅ Public profile fields only
- ✅ Prefix and gender are considered public information

The enhanced responses provide better user experience with complete localization support (Thai/English) while maintaining appropriate privacy levels.