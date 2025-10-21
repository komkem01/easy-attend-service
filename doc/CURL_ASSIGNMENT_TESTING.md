# Assignment API Testing Guide with cURL

## 🧪 Testing Assignments CRUD API

### Prerequisites
1. Start the server: `go run main.go serve`
2. Get JWT token from login endpoint
3. Have a valid classroom_id

---

## 📋 Test Sequence

### 1. รับ JWT Token
```bash
# Login เพื่อรับ token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teacher@example.com",
    "password": "password123"
  }'
```

**Save the token from response:**
```bash
export JWT_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

### 2. ดูรายการ Classrooms (เพื่อเอา classroom_id)
```bash
curl -X GET http://localhost:8080/api/v1/classrooms \
  -H "Authorization: Bearer $JWT_TOKEN"
```

**Save classroom_id from response:**
```bash
export CLASSROOM_ID="123e4567-e89b-12d3-a456-426614174001"
```

---

### 3. สร้าง Assignment แรก
```bash
curl -X POST http://localhost:8080/api/v1/assignments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "classroom_id": "'$CLASSROOM_ID'",
    "title": "โปรแกรมคำนวณเกรด",
    "description": "สร้างโปรแกรมคำนวณเกรดเฉลี่ยนักเรียน",
    "instructions": "ใช้ภาษา Python สร้างโปรแกรมที่สามารถ:\n1. รับคะแนนสอบจากผู้ใช้\n2. คำนวณเกรดเฉลี่ย\n3. แสดงผลเกรดตัวอักษร (A, B+, B, C+, C, D+, D, F)",
    "assignment_type": "homework",
    "due_date": "2024-12-31T23:59:59Z",
    "max_score": 100.0,
    "weight": 1.0,
    "allow_late_submission": true,
    "late_penalty_percent": 10.0,
    "submission_format": "both",
    "max_file_size_mb": 10,
    "allowed_file_types": "[\"py\", \"pdf\", \"docx\", \"txt\"]",
    "is_published": false
  }'
```

**Expected Response:**
```json
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": "generated-uuid-here",
    "title": "โปรแกรมคำนวณเกรด",
    "status": "draft",
    "created_at": "2024-01-20T..."
  }
}
```

**Save assignment_id:**
```bash
export ASSIGNMENT_ID="generated-uuid-from-response"
```

---

### 4. ดูข้อมูล Assignment ที่สร้าง
```bash
curl -X GET http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID
```

---

### 5. แก้ไข Assignment
```bash
curl -X PATCH http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "title": "โปรแกรมคำนวณเกรด - Version 2",
    "max_score": 120.0,
    "instructions": "อัปเดตโปรแกรม:\n1. รับคะแนนสอบจากผู้ใช้\n2. คำนวณเกรดเฉลี่ย\n3. แสดงผลเกรดตัวอักษร\n4. เพิ่มระบบบันทึกผลคะแนน",
    "due_date": "2024-12-25T23:59:59Z"
  }'
```

---

### 6. เผยแพร่ Assignment
```bash
curl -X POST http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID/publish \
  -H "Authorization: Bearer $JWT_TOKEN"
```

**Expected Response:**
```json
{
  "status": {
    "code": 200,
    "message": "Success"
  },
  "data": {
    "id": "assignment-id",
    "is_published": true,
    "status": "active",
    "updated_at": "2024-01-20T..."
  }
}
```

---

### 7. สร้าง Assignment เพิ่มเติม (หลายประเภท)

#### Assignment ประเภท Quiz:
```bash
curl -X POST http://localhost:8080/api/v1/assignments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "classroom_id": "'$CLASSROOM_ID'",
    "title": "แบบทดสอบ Python Basics",
    "description": "ทดสอบความรู้พื้นฐาน Python",
    "instructions": "ตอบคำถาม 20 ข้อ เวลา 30 นาที",
    "assignment_type": "quiz",
    "due_date": "2024-12-20T14:00:00Z",
    "max_score": 20.0,
    "weight": 0.5,
    "allow_late_submission": false,
    "submission_format": "text",
    "is_published": true
  }'
```

#### Assignment ประเภท Project:
```bash
curl -X POST http://localhost:8080/api/v1/assignments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "classroom_id": "'$CLASSROOM_ID'",
    "title": "โปรเจคระบบจัดการหอพัก",
    "description": "พัฒนาระบบจัดการหอพักนักศึกษา",
    "instructions": "สร้างระบบที่มีฟีเจอร์:\n1. ลงทะเบียนเข้าพัก\n2. จัดการห้องพัก\n3. คำนวณค่าใช้จ่าย\n4. รายงานสถานะ",
    "assignment_type": "project",
    "due_date": "2024-12-30T23:59:59Z",
    "max_score": 200.0,
    "weight": 3.0,
    "allow_late_submission": true,
    "late_penalty_percent": 5.0,
    "submission_format": "both",
    "max_file_size_mb": 50,
    "allowed_file_types": "[\"py\", \"js\", \"html\", \"css\", \"pdf\", \"zip\", \"rar\"]",
    "is_published": true
  }'
```

---

### 8. ดูรายการ Assignments (Public)
```bash
# ดูทั้งหมด
curl -X GET http://localhost:8080/api/v1/assignments

# ดูเฉพาะห้องเรียน
curl -X GET "http://localhost:8080/api/v1/assignments?classroom_id=$CLASSROOM_ID"

# ดูเฉพาะที่เผยแพร่แล้ว
curl -X GET "http://localhost:8080/api/v1/assignments?is_published=true"

# ค้นหาด้วยชื่อ
curl -X GET "http://localhost:8080/api/v1/assignments?search=เกรด"

# ดูเฉพาะประเภท homework
curl -X GET "http://localhost:8080/api/v1/assignments?assignment_type=homework"

# ดูงานที่ใกล้ครบกำหนด
curl -X GET "http://localhost:8080/api/v1/assignments?due_soon=true"

# ดูแบบมี pagination
curl -X GET "http://localhost:8080/api/v1/assignments?page=1&limit=5"
```

---

### 9. ทดสอบ Error Cases

#### ทดสอบสร้าง Assignment โดยไม่มี classroom access:
```bash
# ใช้ classroom_id ที่ไม่มีสิทธิ์
curl -X POST http://localhost:8080/api/v1/assignments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "classroom_id": "00000000-0000-0000-0000-000000000000",
    "title": "Test Forbidden",
    "assignment_type": "homework",
    "max_score": 100.0
  }'
```
**Expected: 403 Forbidden**

#### ทดสอบแก้ไขโดยไม่มี token:
```bash
curl -X PATCH http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID \
  -H "Content-Type: application/json" \
  -d '{"title": "Hacked"}'
```
**Expected: 401 Unauthorized**

#### ทดสอบดู Assignment ที่ไม่มี:
```bash
curl -X GET http://localhost:8080/api/v1/assignments/00000000-0000-0000-0000-000000000000
```
**Expected: 404 Not Found**

---

### 10. ลบ Assignment (Soft Delete)
```bash
curl -X DELETE http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID \
  -H "Authorization: Bearer $JWT_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Assignment deleted successfully"
}
```

#### ตรวจสอบว่าลบแล้ว:
```bash
curl -X GET http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID
```
**Expected: 404 Not Found**

---

## 🎯 Complete Test Script

สร้างไฟล์ `test_assignments.sh`:

```bash
#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🧪 Testing Assignment API...${NC}"

# 1. Login
echo -e "\n${YELLOW}1. Getting JWT Token...${NC}"
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "teacher@example.com",
    "password": "password123"
  }')

JWT_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.data.access_token')
echo "Token: ${JWT_TOKEN:0:50}..."

# 2. Get Classrooms
echo -e "\n${YELLOW}2. Getting Classrooms...${NC}"
CLASSROOMS_RESPONSE=$(curl -s -X GET http://localhost:8080/api/v1/classrooms \
  -H "Authorization: Bearer $JWT_TOKEN")

CLASSROOM_ID=$(echo $CLASSROOMS_RESPONSE | jq -r '.data.classrooms[0].id')
echo "Classroom ID: $CLASSROOM_ID"

# 3. Create Assignment
echo -e "\n${YELLOW}3. Creating Assignment...${NC}"
CREATE_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/assignments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "classroom_id": "'$CLASSROOM_ID'",
    "title": "Test Assignment - Auto Generated",
    "description": "Auto generated test assignment",
    "assignment_type": "homework",
    "due_date": "2024-12-31T23:59:59Z",
    "max_score": 100.0,
    "is_published": true
  }')

ASSIGNMENT_ID=$(echo $CREATE_RESPONSE | jq -r '.data.id')
echo "Assignment ID: $ASSIGNMENT_ID"

# 4. Get Assignment
echo -e "\n${YELLOW}4. Getting Assignment Details...${NC}"
curl -s -X GET http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID | jq .

# 5. Update Assignment
echo -e "\n${YELLOW}5. Updating Assignment...${NC}"
curl -s -X PATCH http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{"title": "Updated Test Assignment", "max_score": 120.0}' | jq .

# 6. Publish Assignment
echo -e "\n${YELLOW}6. Publishing Assignment...${NC}"
curl -s -X POST http://localhost:8080/api/v1/assignments/$ASSIGNMENT_ID/publish \
  -H "Authorization: Bearer $JWT_TOKEN" | jq .

# 7. List Assignments
echo -e "\n${YELLOW}7. Listing All Assignments...${NC}"
curl -s -X GET "http://localhost:8080/api/v1/assignments?limit=3" | jq .

echo -e "\n${GREEN}✅ Assignment API Testing Complete!${NC}"
echo -e "Assignment ID for manual testing: ${ASSIGNMENT_ID}"
```

**Run the test:**
```bash
chmod +x test_assignments.sh
./test_assignments.sh
```

---

## 🔍 Expected Results Summary

✅ **Create Assignment**: Status 201, returns assignment ID  
✅ **Get Assignment**: Status 200, returns full assignment data  
✅ **Update Assignment**: Status 200, returns updated fields  
✅ **Publish Assignment**: Status 200, is_published = true  
✅ **List Assignments**: Status 200, returns array with pagination  
✅ **Delete Assignment**: Status 200, success = true  
❌ **Unauthorized Access**: Status 401  
❌ **Forbidden Access**: Status 403  
❌ **Not Found**: Status 404  

Assignment API พร้อมใช้งานและทดสอบครบทุกฟีเจอร์! 🎉