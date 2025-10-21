# Classroom API Testing with cURL

## 1. เริ่มเซิร์ฟเวอร์ก่อน
```bash
# ใช้ air สำหรับ hot reload
air

# หรือใช้ go run
go run main.go serve
```

## 2. ทดสอบ Login เพื่อรับ JWT Token
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teacher@example.com",
    "password": "password123"
  }'
```

**หมายเหตุ:** คุณต้อง register user ก่อน หรือใช้ข้อมูล user ที่มีอยู่แล้ว

## 3. สร้าง Classroom ใหม่ (Protected - ต้องใช้ JWT Token)
```bash
curl -X POST http://localhost:8080/api/v1/classrooms \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{
    "school_id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "คอมพิวเตอร์ 102",
    "subject": "วิทยาการคอมพิวเตอร์",
    "description": "โปรแกรมมิ่งเว็บไซต์",
    "grade_level": "ปวช.2",
    "section": "B",
    "room_number": "CR-102",
    "max_students": 40,
    "schedule": "{\"tuesday\": \"13:00-16:00\"}"
  }'
```

## 4. ดูรายการ Classrooms ทั้งหมด (Public - ไม่ต้อง Token)
```bash
curl -X GET http://localhost:8080/api/v1/classrooms
```

## 5. ดูรายการ Classrooms แบบมี Pagination และ Search
```bash
# แบบมี pagination
curl -X GET "http://localhost:8080/api/v1/classrooms?page=1&limit=10"

# แบบค้นหา
curl -X GET "http://localhost:8080/api/v1/classrooms?search=คอมพิวเตอร์"

# แบบกรองตาม subject
curl -X GET "http://localhost:8080/api/v1/classrooms?subject=วิทยาการคอมพิวเตอร์"
```

## 6. ดู Classroom ตาม ID
```bash
curl -X GET http://localhost:8080/api/v1/classrooms/CLASSROOM_ID_HERE
```

## 7. ค้นหา Classroom ด้วยรหัส
```bash
curl -X GET http://localhost:8080/api/v1/classrooms/code/ABC123
```

## 8. แก้ไข Classroom (Protected - ต้องใช้ JWT Token)
```bash
curl -X PATCH http://localhost:8080/api/v1/classrooms/CLASSROOM_ID_HERE \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{
    "name": "คอมพิวเตอร์ 102 - Advanced",
    "max_students": 30,
    "is_active": true
  }'
```

## 9. ลบ Classroom (Protected - ต้องใช้ JWT Token)
```bash
curl -X DELETE http://localhost:8080/api/v1/classrooms/CLASSROOM_ID_HERE \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## วิธีการใช้งาน Step by Step:

### Step 1: Register หรือ Login
```bash
# Register user ใหม่
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "teacher01",
    "email": "teacher@example.com",
    "password": "password123",
    "first_name": "สมชาย",
    "last_name": "ครูดี",
    "gender": "ชาย",
    "prefix": "นาย",
    "role": "teacher"
  }'

# Login เพื่อรับ JWT Token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teacher@example.com",
    "password": "password123"
  }'
```

### Step 2: Copy JWT Token จาก response
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Step 3: ใช้ Token ในการสร้าง Classroom
```bash
curl -X POST http://localhost:8080/api/v1/classrooms \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{
    "name": "คอมพิวเตอร์ 102",
    "subject": "วิทยาการคอมพิวเตอร์",
    "description": "โปรแกรมมิ่งเว็บไซต์",
    "grade_level": "ปวช.2",
    "section": "B",
    "room_number": "CR-102",
    "max_students": 40
  }'
```

## Expected Responses:

### สำเร็จ (201 Created):
```json
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "name": "คอมพิวเตอร์ 102",
    "subject": "วิทยาการคอมพิวเตอร์",
    "classroom_code": "XYZ789",
    "created_at": "2024-01-20T15:30:00Z"
  }
}
```

### Error (401 Unauthorized):
```json
{
  "code": 401,
  "message": "User not authenticated"
}
```

### Error (400 Bad Request):
```json
{
  "code": 400,
  "message": "Invalid request data: name is required"
}
```

## หมายเหตุ:
- แทนที่ `YOUR_JWT_TOKEN_HERE` ด้วย JWT token จริงที่ได้จากการ login
- แทนที่ `CLASSROOM_ID_HERE` ด้วย UUID ของ classroom จริง
- เซิร์ฟเวอร์ต้องรันอยู่บนพอร์ต 8080