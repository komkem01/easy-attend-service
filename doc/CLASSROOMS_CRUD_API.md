# Classrooms CRUD API Documentation

## ภาพรวม
API สำหรับการจัดการห้องเรียน (Classrooms) ในระบบ Easy Attend Service

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
การจัดการห้องเรียน (POST, PATCH, DELETE) ต้องการ JWT token ใน Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## 📚 Endpoints

### 1. ดูรายการห้องเรียน (Public)
```http
GET /classrooms
```

#### Query Parameters:
- `page` (optional): หมายเลขหน้า (default: 1)
- `limit` (optional): จำนวนรายการต่อหน้า (default: 20, max: 100)
- `search` (optional): ค้นหาในชื่อห้อง, วิชา, หรือรหัสห้อง
- `school_id` (optional): UUID ของโรงเรียน
- `teacher_id` (optional): UUID ของครู
- `subject` (optional): ชื่อวิชา
- `grade_level` (optional): ระดับชั้น
- `is_active` (optional): สถานะการใช้งาน (true/false)

#### Response:
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "classrooms": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "school_id": "123e4567-e89b-12d3-a456-426614174001",
        "name": "คอมพิวเตอร์ 101",
        "subject": "วิทยาการคอมพิวเตอร์",
        "description": "หลักการโปรแกรมมิ่งเบื้องต้น",
        "grade_level": "ปวช.1",
        "section": "A",
        "room_number": "CR-101",
        "teacher_id": "123e4567-e89b-12d3-a456-426614174002",
        "classroom_code": "ABC123",
        "max_students": 50,
        "schedule": "{\"monday\": \"09:00-12:00\"}",
        "is_active": true,
        "created_at": "2024-01-20T10:00:00Z",
        "updated_at": "2024-01-20T10:00:00Z",
        "school": {
          "id": "123e4567-e89b-12d3-a456-426614174001",
          "name": "โรงเรียนเทคโนโลยี"
        },
        "teacher": {
          "id": "123e4567-e89b-12d3-a456-426614174002",
          "first_name": "สมชาย",
          "last_name": "ใจดี"
        }
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total": 150,
      "total_pages": 8,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

### 2. ดูข้อมูลห้องเรียนตาม ID (Public)
```http
GET /classrooms/{id}
```

#### Response:
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "คอมพิวเตอร์ 101",
    "subject": "วิทยาการคอมพิวเตอร์",
    "classroom_code": "ABC123",
    "teacher": {
      "first_name": "สมชาย",
      "last_name": "ใจดี"
    },
    "classroom_students": [
      {
        "student": {
          "first_name": "นางสาวจิรา",
          "last_name": "เก่งเรียน"
        }
      }
    ]
  }
}
```

### 3. ค้นหาห้องเรียนด้วยรหัส (Public)
```http
GET /classrooms/code/{code}
```

#### Response:
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "คอมพิวเตอร์ 101",
    "classroom_code": "ABC123",
    "teacher": {
      "first_name": "สมชาย",
      "last_name": "ใจดี"
    }
  }
}
```

### 4. สร้างห้องเรียนใหม่ (Protected)
```http
POST /classrooms
```

#### Headers:
```
Authorization: Bearer <jwt-token>
Content-Type: application/json
```

#### Request Body:
```json
{
  "school_id": "123e4567-e89b-12d3-a456-426614174001",
  "name": "คอมพิวเตอร์ 102",
  "subject": "วิทยาการคอมพิวเตอร์",
  "description": "โปรแกรมมิ่งเว็บไซต์",
  "grade_level": "ปวช.2",
  "section": "B",
  "room_number": "CR-102",
  "max_students": 40,
  "schedule": "{\"tuesday\": \"13:00-16:00\"}"
}
```

#### Response:
```json
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174003",
    "name": "คอมพิวเตอร์ 102",
    "classroom_code": "XYZ789",
    "created_at": "2024-01-20T15:30:00Z"
  }
}
```

### 5. แก้ไขห้องเรียน (Protected)
```http
PATCH /classrooms/{id}
```

#### Headers:
```
Authorization: Bearer <jwt-token>
Content-Type: application/json
```

#### Request Body (ส่งเฉพาะฟิลด์ที่ต้องการแก้ไข):
```json
{
  "name": "คอมพิวเตอร์ 102 - Advanced",
  "max_students": 30,
  "is_active": true
}
```

#### Response:
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174003",
    "name": "คอมพิวเตอร์ 102 - Advanced",
    "max_students": 30,
    "updated_at": "2024-01-20T16:00:00Z"
  }
}
```

### 6. ลบห้องเรียน (Protected)
```http
DELETE /classrooms/{id}
```

#### Headers:
```
Authorization: Bearer <jwt-token>
```

#### Response:
```json
{
  "success": true,
  "message": "Classroom deleted successfully"
}
```

---

## 🚨 Error Responses

### 400 Bad Request
```json
{
  "code": 400,
  "message": "Invalid request data: name is required"
}
```

### 401 Unauthorized
```json
{
  "code": 401,
  "message": "User not authenticated"
}
```

### 403 Forbidden
```json
{
  "code": 403,
  "message": "You don't have permission to update this classroom"
}
```

### 404 Not Found
```json
{
  "code": 404,
  "message": "Classroom not found"
}
```

### 500 Internal Server Error
```json
{
  "code": 500,
  "message": "Failed to create classroom: database error"
}
```

---

## 🔐 Authorization Rules

### Public Access (ไม่ต้อง login)
- ✅ `GET /classrooms` - ดูรายการห้องเรียน
- ✅ `GET /classrooms/{id}` - ดูข้อมูลห้องเรียน
- ✅ `GET /classrooms/code/{code}` - ค้นหาห้องด้วยรหัส

### Protected Access (ต้อง login)
- 🔒 `POST /classrooms` - สร้างห้องเรียน (ครูเท่านั้น)
- 🔒 `PATCH /classrooms/{id}` - แก้ไขห้องเรียน (ครูผู้สอนเท่านั้น)
- 🔒 `DELETE /classrooms/{id}` - ลบห้องเรียน (ครูผู้สอนเท่านั้น)

---

## 💡 Features

### 🔑 **Automatic Classroom Code Generation**
- สร้างรหัสห้องเรียน 6 ตัวอักษร (A-Z, 0-9) อัตโนมัติ
- ตรวจสอบความไม่ซ้ำกัน

### 🛡️ **Soft Delete Protection**
- ไม่สามารถลบห้องที่มีนักเรียนที่ยัง active อยู่
- Soft delete - ข้อมูลไม่หายจริง

### 🔍 **Advanced Search & Filter**
- ค้นหาด้วยชื่อห้อง, วิชา, รหัสห้อง
- กรองตามโรงเรียน, ครู, ระดับชั้น
- Pagination แบบมีประสิทธิภาพ

### 👥 **Relationship Loading**
- โหลดข้อมูลโรงเรียนและครูผู้สอน
- แสดงรายชื่อนักเรียนในห้อง (เฉพาะ GET by ID)

### 🔐 **Permission Control**
- เฉพาะครูผู้สอนสามารถแก้ไข/ลบห้องตนเอง
- Admin สามารถจัดการทุกห้อง (TODO)

ระบบ Classrooms CRUD พร้อมใช้งานแล้ว! 🎉