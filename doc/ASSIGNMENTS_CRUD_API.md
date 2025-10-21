# Assignments CRUD API Documentation

## ภาพรวม
API สำหรับการจัดการงานที่มอบหมาย (Assignments) ในระบบ Easy Attend Service

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
การจัดการงาน (POST, PATCH, DELETE) ต้องการ JWT token ใน Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## 📚 Endpoints

### 1. ดูรายการงานทั้งหมด (Public)
```http
GET /assignments
```

#### Query Parameters:
- `page` (optional): หมายเลขหน้า (default: 1)
- `limit` (optional): จำนวนรายการต่อหน้า (default: 20, max: 100)
- `search` (optional): ค้นหาในชื่องาน, คำอธิบาย
- `classroom_id` (optional): UUID ของห้องเรียน
- `created_by` (optional): UUID ของผู้สร้าง
- `assignment_type` (optional): ประเภทงาน (homework, quiz, exam, project, lab_work)
- `status` (optional): สถานะงาน (draft, active, completed, archived)
- `is_published` (optional): สถานะการเผยแพร่ (true/false)
- `due_soon` (optional): งานที่ใกล้ครบกำหนด (7 วัน) (true/false)
- `overdue` (optional): งานที่เลยกำหนดแล้ว (true/false)

#### Response:
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "assignments": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "classroom_id": "123e4567-e89b-12d3-a456-426614174001",
        "title": "โปรแกรมคำนวณเกรด",
        "description": "สร้างโปรแกรมคำนวณเกรดเฉลี่ย",
        "instructions": "ใช้ภาษา Python สร้างโปรแกรม...",
        "assignment_type": "homework",
        "due_date": "2024-01-30T23:59:59Z",
        "max_score": 100.0,
        "weight": 1.0,
        "allow_late_submission": true,
        "late_penalty_percent": 10.0,
        "submission_format": "both",
        "max_file_size_mb": 10,
        "allowed_file_types": "[\"py\", \"pdf\", \"docx\"]",
        "is_published": true,
        "status": "active",
        "created_by": "123e4567-e89b-12d3-a456-426614174002",
        "created_at": "2024-01-20T10:00:00Z",
        "updated_at": "2024-01-20T10:00:00Z",
        "classroom": {
          "id": "123e4567-e89b-12d3-a456-426614174001",
          "name": "คอมพิวเตอร์ 101",
          "subject": "วิทยาการคอมพิวเตอร์"
        },
        "creator": {
          "id": "123e4567-e89b-12d3-a456-426614174002",
          "first_name": "สมชาย",
          "last_name": "ครูดี"
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

### 2. ดูข้อมูลงานตาม ID (Public)
```http
GET /assignments/{id}
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
    "title": "โปรแกรมคำนวณเกรด",
    "description": "สร้างโปรแกรมคำนวณเกรดเฉลี่ย",
    "instructions": "ใช้ภาษา Python สร้างโปรแกรม...",
    "due_date": "2024-01-30T23:59:59Z",
    "max_score": 100.0,
    "classroom": {
      "name": "คอมพิวเตอร์ 101",
      "teacher": {
        "first_name": "สมชาย",
        "last_name": "ครูดี"
      }
    }
  }
}
```

### 3. สร้างงานใหม่ (Protected)
```http
POST /assignments
```

#### Headers:
```
Authorization: Bearer <jwt-token>
Content-Type: application/json
```

#### Request Body:
```json
{
  "classroom_id": "123e4567-e89b-12d3-a456-426614174001",
  "title": "โปรแกรมคำนวณเกรด",
  "description": "สร้างโปรแกรมคำนวณเกรดเฉลี่ย",
  "instructions": "ใช้ภาษา Python สร้างโปรแกรมที่สามารถ:\n1. รับคะแนนสอบ\n2. คำนวณเกรดเฉลี่ย\n3. แสดงผลเกรด",
  "assignment_type": "homework",
  "due_date": "2024-01-30T23:59:59Z",
  "max_score": 100.0,
  "weight": 1.0,
  "allow_late_submission": true,
  "late_penalty_percent": 10.0,
  "submission_format": "both",
  "max_file_size_mb": 10,
  "allowed_file_types": "[\"py\", \"pdf\", \"docx\"]",
  "is_published": false
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
    "title": "โปรแกรมคำนวณเกรด",
    "status": "draft",
    "created_at": "2024-01-20T15:30:00Z"
  }
}
```

### 4. แก้ไขงาน (Protected)
```http
PATCH /assignments/{id}
```

#### Headers:
```
Authorization: Bearer <jwt-token>
Content-Type: application/json
```

#### Request Body (ส่งเฉพาะฟิลด์ที่ต้องการแก้ไข):
```json
{
  "title": "โปรแกรมคำนวณเกรด - Advanced",
  "due_date": "2024-02-05T23:59:59Z",
  "max_score": 120.0,
  "is_published": true,
  "status": "active"
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
    "title": "โปรแกรมคำนวณเกรด - Advanced",
    "max_score": 120.0,
    "is_published": true,
    "status": "active",
    "updated_at": "2024-01-20T16:00:00Z"
  }
}
```

### 5. ลบงาน (Protected)
```http
DELETE /assignments/{id}
```

#### Headers:
```
Authorization: Bearer <jwt-token>
```

#### Response:
```json
{
  "success": true,
  "message": "Assignment deleted successfully"
}
```

### 6. เผยแพร่งาน (Protected)
```http
POST /assignments/{id}/publish
```

#### Headers:
```
Authorization: Bearer <jwt-token>
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
    "is_published": true,
    "status": "active",
    "updated_at": "2024-01-20T16:30:00Z"
  }
}
```

---

## 📋 Field Descriptions

### Assignment Types
- `homework` - การบ้าน
- `quiz` - แบบทดสอบ
- `exam` - สอบ
- `project` - โปรเจค
- `lab_work` - งานปฏิบัติการ

### Assignment Status
- `draft` - ร่าง (ยังไม่เผยแพร่)
- `active` - ใช้งาน (เผยแพร่แล้ว)
- `completed` - เสร็จสิ้น
- `archived` - เก็บถาวร

### Submission Format
- `text` - ส่งเฉพาะข้อความ
- `file` - ส่งเฉพาะไฟล์
- `both` - ส่งได้ทั้งข้อความและไฟล์

---

## 🚨 Error Responses

### 400 Bad Request
```json
{
  "code": 400,
  "message": "Invalid request data: title is required"
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
  "message": "You don't have access to this classroom"
}
```

### 404 Not Found
```json
{
  "code": 404,
  "message": "Assignment not found"
}
```

### 500 Internal Server Error
```json
{
  "code": 500,
  "message": "Failed to create assignment: database error"
}
```

---

## 🔐 Authorization Rules

### Public Access (ไม่ต้อง login)
- ✅ `GET /assignments` - ดูรายการงาน
- ✅ `GET /assignments/{id}` - ดูข้อมูลงาน

### Protected Access (ต้อง login)
- 🔒 `POST /assignments` - สร้างงาน (ครูในห้องเรียนนั้นเท่านั้น)
- 🔒 `PATCH /assignments/{id}` - แก้ไขงาน (ผู้สร้างหรือครูผู้สอนเท่านั้น)
- 🔒 `DELETE /assignments/{id}` - ลบงาน (ผู้สร้างหรือครูผู้สอนเท่านั้น)
- 🔒 `POST /assignments/{id}/publish` - เผยแพร่งาน (ผู้สร้างหรือครูผู้สอนเท่านั้น)

---

## 💡 Features

### 🎯 **Assignment Management**
- สร้างงานหลากหลายประเภท (การบ้าน, แบบทดสอบ, โปรเจค)
- กำหนดคะแนนเต็ม, น้ำหนัก, วันที่ครบกำหนด
- ตั้งค่าการส่งงานช้า และการหักคะแนน

### 📁 **File Submission Control**
- กำหนดรูปแบบการส่งงาน (ข้อความ, ไฟล์, หรือทั้งคู่)
- จำกัดขนาดไฟล์และประเภทไฟล์ที่อนุญาต
- รองรับไฟล์หลากหลายรูปแบบ

### 🔍 **Advanced Search & Filter**
- ค้นหาด้วยชื่องาน, คำอธิบาย
- กรองตามห้องเรียน, ประเภทงาน, สถานะ
- ดูงานที่ใกล้ครบกำหนดหรือเลยกำหนด

### 📊 **Status Management**
- Draft → Active → Completed → Archived
- เผยแพร่งานเมื่อพร้อม
- ซ่อนงานร่างจากนักเรียน

### 🛡️ **Permission Control**
- เฉพาะครูในห้องเรียนสร้างงานได้
- ผู้สร้างและครูผู้สอนแก้ไข/ลบได้
- ป้องกันการลบงานที่มีการส่งแล้ว

### 🔄 **Soft Delete Protection**
- ไม่สามารถลบงานที่มีการส่งแล้ว
- ข้อมูลไม่สูญหายจริง

## cURL Testing Examples

### สร้าง Assignment:
```bash
curl -X POST http://localhost:8080/api/v1/assignments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "classroom_id": "123e4567-e89b-12d3-a456-426614174001",
    "title": "โปรแกรมคำนวณเกรด",
    "description": "สร้างโปรแกรมคำนวณเกรดเฉลี่ย",
    "assignment_type": "homework",
    "due_date": "2024-01-30T23:59:59Z",
    "max_score": 100.0,
    "allow_late_submission": true,
    "late_penalty_percent": 10.0
  }'
```

### ดูรายการ Assignments:
```bash
curl -X GET "http://localhost:8080/api/v1/assignments?classroom_id=123e4567-e89b-12d3-a456-426614174001&is_published=true"
```

### เผยแพร่ Assignment:
```bash
curl -X POST http://localhost:8080/api/v1/assignments/YOUR_ASSIGNMENT_ID/publish \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

ระบบ Assignments CRUD พร้อมใช้งานครบทุกฟีเจอร์แล้ว! 🎉