# Air Hot Reload สำหรับ Easy Attend Service

## การติดตั้งและใช้งาน

### 1. ติดตั้ง Air ✅
```powershell
go install github.com/air-verse/air@latest
```

### 2. สร้างไฟล์ config ✅
```powershell
air init
```

### 3. การใช้งาน
```powershell
# รันโปรเจคด้วย hot reload
air

# หรือใช้คำสั่งเต็ม
air serve
```

## การตั้งค่า (.air.toml)

### Build Configuration
- **Binary**: `tmp/easy-attend.exe`
- **Command**: `go build -o ./tmp/easy-attend.exe .`
- **Args**: `["serve"]` - จะรัน `serve` command อัตโนมัติ
- **Delay**: 1000ms - รอก่อน restart

### File Watching
- **Include**: `.go`, `.mod`, `.env` files
- **Exclude**: 
  - Directories: `assets`, `tmp`, `vendor`, `testdata`
  - Files: `*_test.go`

### Hot Reload Features
- 🔥 **Auto Restart**: เมื่อไฟล์ .go เปลี่ยน
- 📦 **Auto Build**: Compile โปรเจคใหม่อัตโนมัติ
- 🚀 **Quick Start**: รันเซิร์ฟเวอร์ทันที
- 📝 **Build Logs**: บันทึกข้อผิดพลาดใน `build-errors.log`

## วิธีใช้งาน

### เริ่มต้น Development Server
```powershell
cd c:\Users\kemke\easy-attend-service
air
```

### ผลลัพธ์ที่คาดหวัง
```
  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.63.0

watching .
!exclude tmp
building...
running...

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
[GIN-debug] GET    /api/v1/health            --> main.main.func1
[GIN-debug] POST   /api/v1/auth/register     --> ...
Server starting on :8080
```

### การทำงานของ Hot Reload
1. แก้ไขไฟล์ `.go` ใดๆ
2. Air จะตรวจจับการเปลี่ยนแปลง
3. Build โปรเจคใหม่อัตโนมัติ
4. Restart server ด้วยโค้ดใหม่
5. API พร้อมใช้งานทันที

### หยุดการทำงาน
```powershell
Ctrl + C
```

## ข้อดีของ Air
- ⚡ **เร็ว**: Build และ restart เร็วกว่า manual
- 🔄 **อัตโนมัติ**: ไม่ต้องหยุด-เริ่มเอง
- 🛠️ **Flexible**: ปรับแต่ง config ได้
- 🎯 **Focus**: เขียนโค้ดโดยไม่ต้องกังวลเรื่อง restart

พร้อมใช้งาน Development ด้วย Hot Reload แล้ว! 🎉