# Soft Delete Implementation - เพิ่ม deleted_at ให้ทุกตาราง

## สรุปการเปลี่ยนแปลง ✅

### ตารางที่เพิ่ม deleted_at แล้ว:

#### 1. ตารางหลัก (Core Tables)
- ✅ **users** - ผู้ใช้งาน
- ✅ **schools** - โรงเรียน  
- ✅ **genders** - เพศ
- ✅ **prefixes** - คำนำหน้าชื่อ

#### 2. ตารางระบบ (System Tables)
- ✅ **system_settings** - การตั้งค่าระบบ
- ✅ **session_tokens** - โทเค็น session
- ✅ **user_sessions** - เซสชันผู้ใช้
- ✅ **user_role_permissions** - สิทธิ์ผู้ใช้
- ✅ **user_profiles** - โปรไฟล์ผู้ใช้

#### 3. ตารางการสื่อสารและบันทึก (Communication & Logs)
- ✅ **notifications** - การแจ้งเตือน
- ✅ **messages** - ข้อความ (แก้ไข soft_delete tag)
- ✅ **security_events** - เหตุการณ์ความปลอดภัย
- ✅ **search_logs** - บันทึกการค้นหา
- ✅ **metrics_data** - ข้อมูลเมตริก

#### 4. ตารางไฟล์ (File Management)
- ✅ **file_uploads** - การอัปโหลดไฟล์ (แก้ไข soft_delete tag)

#### 5. ตารางที่มี deleted_at อยู่แล้ว
- ✅ **classrooms** - ห้องเรียน (มีอยู่แล้ว)

## รูปแบบ deleted_at ที่ใช้

### Standard Format
```go
DeletedAt *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`
```

### คุณสมบัติ:
- **Type**: `*time.Time` (nullable)
- **JSON**: `omitempty` - ไม่แสดงใน JSON เมื่อ null
- **Bun ORM**: `soft_delete` - เปิดใช้งาน soft delete อัตโนมัติ

## การทำงานของ Soft Delete

### 1. การลบ (Soft Delete)
```go
// แทนที่จะลบจริง จะเซ็ต deleted_at = now()
db.NewDelete().Model(&user).Where("id = ?", userID).Exec(ctx)
```

### 2. การดึงข้อมูล (Auto Filter)
```go
// จะไม่แสดงข้อมูลที่มี deleted_at != null อัตโนมัติ  
db.NewSelect().Model(&users).Scan(ctx, &users)
```

### 3. การดึงข้อมูลรวมที่ถูกลบ
```go
// ดึงข้อมูลทั้งหมดรวมที่ถูกลบ
db.NewSelect().Model(&users).WhereDeleted().Scan(ctx, &users)
```

### 4. การกู้คืนข้อมูล (Restore)
```go
// เซ็ต deleted_at = null
db.NewUpdate().Model(&user).
    Set("deleted_at = NULL").
    Where("id = ?", userID).Exec(ctx)
```

## Database Migration

### ไฟล์ Migration: `20241021_add_deleted_at_columns.go`

#### Up Migration:
- ตรวจสอบว่าคอลัมน์ `deleted_at` มีอยู่แล้วหรือไม่
- เพิ่มคอลัมน์ `deleted_at TIMESTAMP DEFAULT NULL` 
- ใช้กับ 13 ตาราง

#### Down Migration:
- ลบคอลัมน์ `deleted_at` ออกจากทุกตาราง

### วิธีรัน Migration:
```powershell
# เพิ่ม deleted_at columns
go run main.go migrate

# หรือใช้ air (จะ migrate อัตโนมัติ)
air
```

## ประโยชน์ของ Soft Delete

### 1. ความปลอดภัยข้อมูล 🛡️
- ป้องกันการลบข้อมูลโดยไม่ตั้งใจ
- สามารถกู้คืนข้อมูลได้

### 2. การตรวจสอบย้อนหลัง 📊
- ติดตามประวัติการลบ
- วิเคราะห์พฤติกรรมผู้ใช้

### 3. ความสัมพันธ์ข้อมูล 🔗
- ไม่ทำลายความสัมพันธ์ระหว่างตาราง
- Foreign key constraints ยังคงทำงาน

### 4. การปฏิบัติตาม Regulations 📋
- GDPR compliance
- ข้อกำหนดการเก็บข้อมูล

## ข้อควรระวัง ⚠️

### 1. ขนาดฐานข้อมูล
- ข้อมูลที่ถูก soft delete ยังคงอยู่
- ต้องมีกระบวนการทำความสะอาดข้อมูลเก่า

### 2. Performance
- Query อาจช้าลงเล็กน้อยเนื่องจาก WHERE deleted_at IS NULL
- ควรสร้าง index บน deleted_at column

### 3. Business Logic
- ต้องระวังเรื่อง unique constraints
- อาจต้องเพิ่มเงื่อนไข deleted_at ใน unique index

## สถานะปัจจุบัน

✅ **พร้อมใช้งาน**: ทุกตารางสำคัญมี soft delete แล้ว  
✅ **Migration พร้อม**: ไฟล์ migration สร้างเสร็จแล้ว  
✅ **ORM Support**: Bun ORM รองรับ soft delete อัตโนมัติ  

ระบบ Soft Delete ได้ถูกเพิ่มให้กับทุกตารางเรียบร้อยแล้ว! 🎉