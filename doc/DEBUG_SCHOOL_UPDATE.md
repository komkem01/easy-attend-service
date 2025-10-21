# Test School Update Issue

## Problem
When updating school information using PATCH method, the system returns "school name already exists" error even when updating the same school.

## Debug Steps

### 1. Start the server
```powershell
.\easy-attend.exe serve
```

### 2. Create a test school first
```powershell
# Login first to get token
$loginData = @{
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method POST -Body $loginData -ContentType "application/json"
$token = $loginResponse.data.access_token

# Create a test school
$schoolData = @{
    name = "Test University"
    address = "123 Test Street"
    phone = "02-123-4567"
    email = "info@testuni.ac.th"
    website_url = "https://testuni.ac.th"
} | ConvertTo-Json

$headers = @{ 
    Authorization = "Bearer $token"
    'Content-Type' = 'application/json'
}

$school = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/schools" -Method POST -Body $schoolData -Headers $headers
$schoolId = $school.data.id
Write-Host "Created school with ID: $schoolId"
```

### 3. Try to update the school (this should trigger the bug)
```powershell
# Update with SAME name (should work)
$updateData = @{
    name = "Test University"
    phone = "02-123-4568"
} | ConvertTo-Json

try {
    $updateResult = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/schools/$schoolId" -Method PATCH -Body $updateData -Headers $headers
    Write-Host "Update successful!" -ForegroundColor Green
    $updateResult | ConvertTo-Json -Depth 5
} catch {
    Write-Host "Update failed with error:" -ForegroundColor Red
    $_.Exception.Message
    if ($_.ErrorDetails) {
        $_.ErrorDetails.Message
    }
}
```

### 4. Update with DIFFERENT name (should also work)
```powershell
$updateData2 = @{
    name = "Test University Updated"
    phone = "02-123-4569"
} | ConvertTo-Json

try {
    $updateResult2 = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/schools/$schoolId" -Method PATCH -Body $updateData2 -Headers $headers
    Write-Host "Update with new name successful!" -ForegroundColor Green
    $updateResult2 | ConvertTo-Json -Depth 5
} catch {
    Write-Host "Update with new name failed:" -ForegroundColor Red
    $_.Exception.Message
    if ($_.ErrorDetails) {
        $_.ErrorDetails.Message
    }
}
```

## Expected Debug Output
The server should show debug messages like:
```
DEBUG: Current name: 'Test University', New name: 'Test University', School ID: [uuid]
DEBUG: Names are the same, skipping duplicate check.
```

## Possible Issues

1. **String comparison issue**: Current name vs new name might have whitespace differences
2. **Case sensitivity**: LOWER() function might not work as expected
3. **UUID comparison**: The WHERE clause might not properly exclude the current school
4. **Transaction issue**: Update might be running in a transaction that sees stale data

## Solution Applied

1. Added debug logging to see exact values being compared
2. Added explicit check to compare current name vs new name before doing duplicate check
3. Only run duplicate check if names are actually different

## Alternative Solution (if debug shows issues)

If the issue persists, we can:

1. **Remove duplicate check for updates**: Only check duplicates on creation
2. **Use username validation pattern**: Check if current user owns the school being updated
3. **Add more sophisticated name comparison**: Trim whitespace, normalize strings

```go
// Alternative implementation without duplicate check on updates
func UpdateSchoolService(ctx context.Context, schoolID string, updateData map[string]interface{}) (*model.Schools, error) {
    // Skip duplicate name check for updates entirely
    // Just update and let database constraints handle it
    
    // ... rest of update logic
}
```