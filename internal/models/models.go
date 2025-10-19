package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Custom types for enums
type UserRole string
type AttendanceStatus string
type SessionStatus string
type SessionMethod string
type CheckInMethod string
type MessageType string
type NotificationType string
type PriorityLevel string
type AssignmentType string
type SubmissionFormat string
type AssignmentStatus string
type SubmissionStatus string
type FileCategory string
type PermissionType string
type EventType string
type RiskLevel string
type MetricType string
type SearchType string
type DevicePlatform string
type DataTypeSetting string
type ReferenceType string

// Enum constants
const (
	// User roles
	RoleStudent UserRole = "student"
	RoleTeacher UserRole = "teacher"
	RoleAdmin   UserRole = "admin"

	// Attendance status
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
	StatusExcused AttendanceStatus = "excused"

	// Session status
	SessionScheduled SessionStatus = "scheduled"
	SessionActive    SessionStatus = "active"
	SessionEnded     SessionStatus = "ended"
	SessionCancelled SessionStatus = "cancelled"

	// Session methods
	MethodCode   SessionMethod = "code"
	MethodQR     SessionMethod = "qr"
	MethodManual SessionMethod = "manual"
	MethodAuto   SessionMethod = "auto"

	// Check-in methods
	CheckInCode   CheckInMethod = "code"
	CheckInQR     CheckInMethod = "qr"
	CheckInManual CheckInMethod = "manual"
	CheckInAuto   CheckInMethod = "auto"

	// Message types
	MessagePrivate      MessageType = "private"
	MessageClassroom    MessageType = "classroom"
	MessageAnnouncement MessageType = "announcement"

	// Notification types
	NotificationAttendance   NotificationType = "attendance"
	NotificationAssignment   NotificationType = "assignment"
	NotificationGrade        NotificationType = "grade"
	NotificationAnnouncement NotificationType = "announcement"
	NotificationSystem       NotificationType = "system"

	// Priority levels
	PriorityLow    PriorityLevel = "low"
	PriorityNormal PriorityLevel = "normal"
	PriorityHigh   PriorityLevel = "high"
	PriorityUrgent PriorityLevel = "urgent"

	// Assignment types
	AssignmentHomework AssignmentType = "homework"
	AssignmentProject  AssignmentType = "project"
	AssignmentQuiz     AssignmentType = "quiz"
	AssignmentExam     AssignmentType = "exam"
	AssignmentLab      AssignmentType = "lab"

	// Submission formats
	SubmissionFile SubmissionFormat = "file"
	SubmissionText SubmissionFormat = "text"
	SubmissionBoth SubmissionFormat = "both"

	// Assignment status
	AssignmentDraft     AssignmentStatus = "draft"
	AssignmentPublished AssignmentStatus = "published"
	AssignmentClosed    AssignmentStatus = "closed"

	// Submission status
	SubmissionDraft     SubmissionStatus = "draft"
	SubmissionSubmitted SubmissionStatus = "submitted"
	SubmissionGraded    SubmissionStatus = "graded"
	SubmissionReturned  SubmissionStatus = "returned"

	// File categories
	FileCategoryAvatar     FileCategory = "avatar"
	FileCategoryAssignment FileCategory = "assignment"
	FileCategorySubmission FileCategory = "submission"
	FileCategoryClassroom  FileCategory = "classroom"
	FileCategoryOther      FileCategory = "other"

	// Permission types
	PermissionView     PermissionType = "view"
	PermissionDownload PermissionType = "download"
	PermissionEdit     PermissionType = "edit"
	PermissionDelete   PermissionType = "delete"

	// Event types
	EventHoliday       EventType = "holiday"
	EventExam          EventType = "exam"
	EventBreak         EventType = "break"
	EventSemesterStart EventType = "semester_start"
	EventSemesterEnd   EventType = "semester_end"
	EventOther         EventType = "other"

	// Risk levels
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"

	// Metric types
	MetricAttendanceRate  MetricType = "attendance_rate"
	MetricAssignmentScore MetricType = "assignment_score"
	MetricExamScore       MetricType = "exam_score"
	MetricParticipation   MetricType = "participation"

	// Search types
	SearchStudents    SearchType = "students"
	SearchAssignments SearchType = "assignments"
	SearchAttendance  SearchType = "attendance"
	SearchClassrooms  SearchType = "classrooms"
	SearchGeneral     SearchType = "general"

	// Device platforms
	PlatformIOS     DevicePlatform = "ios"
	PlatformAndroid DevicePlatform = "android"
	PlatformWeb     DevicePlatform = "web"

	// Data type settings
	DataTypeString  DataTypeSetting = "string"
	DataTypeNumber  DataTypeSetting = "number"
	DataTypeBoolean DataTypeSetting = "boolean"
	DataTypeJSON    DataTypeSetting = "json"

	// Reference types
	ReferenceAttendanceSession ReferenceType = "attendance_session"
	ReferenceAssignment        ReferenceType = "assignment"
	ReferenceClassroom         ReferenceType = "classroom"
)

// Base model with UUID and common fields
type BaseModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// School represents a school/institution
type School struct {
	BaseModel
	Name       string `json:"name" gorm:"not null"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	WebsiteURL string `json:"website_url"`
	LogoURL    string `json:"logo_url"`
	IsActive   bool   `json:"is_active" gorm:"default:true"`
}

// User represents a user in the system
type User struct {
	BaseModel
	SchoolID      *uuid.UUID `json:"school_id"`
	Email         string     `json:"email" gorm:"unique;not null"`
	PasswordHash  string     `json:"-" gorm:"column:password_hash;not null"`
	Name          string     `json:"name" gorm:"not null"`
	Role          UserRole   `json:"role" gorm:"type:user_role;not null"`
	AvatarURL     string     `json:"avatar_url"`
	Phone         string     `json:"phone"`
	DateOfBirth   *time.Time `json:"date_of_birth" gorm:"type:date"`
	Address       string     `json:"address"`
	IsActive      bool       `json:"is_active" gorm:"default:true"`
	EmailVerified bool       `json:"email_verified" gorm:"default:false"`
	LastLoginAt   *time.Time `json:"last_login_at"`

	// Relations
	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

// Classroom represents a classroom/course
type Classroom struct {
	BaseModel
	SchoolID      *uuid.UUID `json:"school_id"`
	Name          string     `json:"name" gorm:"not null"`
	Subject       string     `json:"subject" gorm:"not null"`
	Description   string     `json:"description"`
	GradeLevel    string     `json:"grade_level"`
	Section       string     `json:"section"`
	RoomNumber    string     `json:"room_number"`
	TeacherID     uuid.UUID  `json:"teacher_id" gorm:"not null"`
	ClassroomCode string     `json:"classroom_code" gorm:"unique;not null"`
	MaxStudents   int        `json:"max_students" gorm:"default:50"`
	Schedule      string     `json:"schedule" gorm:"type:jsonb"`
	IsActive      bool       `json:"is_active" gorm:"default:true"`

	// Relations
	School  *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	Teacher *User   `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`
}

// ClassroomStudent represents the enrollment of a student in a classroom
type ClassroomStudent struct {
	BaseModel
	ClassroomID   uuid.UUID `json:"classroom_id" gorm:"not null"`
	StudentID     uuid.UUID `json:"student_id" gorm:"not null"`
	StudentNumber string    `json:"student_number"`
	SeatNumber    string    `json:"seat_number"`
	EnrolledAt    time.Time `json:"enrolled_at" gorm:"default:now()"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	Notes         string    `json:"notes"`

	// Relations
	Classroom *Classroom `json:"classroom,omitempty" gorm:"foreignKey:ClassroomID"`
	Student   *User      `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

// AttendanceSession represents an attendance session
type AttendanceSession struct {
	BaseModel
	ClassroomID          uuid.UUID     `json:"classroom_id" gorm:"not null"`
	Title                string        `json:"title" gorm:"not null"`
	Description          string        `json:"description"`
	SessionDate          time.Time     `json:"session_date" gorm:"type:date;not null"`
	StartTime            time.Time     `json:"start_time" gorm:"type:time;not null"`
	EndTime              time.Time     `json:"end_time" gorm:"type:time;not null"`
	ActualStartTime      *time.Time    `json:"actual_start_time"`
	ActualEndTime        *time.Time    `json:"actual_end_time"`
	Status               SessionStatus `json:"status" gorm:"type:session_status;default:scheduled"`
	Method               SessionMethod `json:"method" gorm:"type:session_method;default:code"`
	SessionCode          string        `json:"session_code"`
	QRCodeData           string        `json:"qr_code_data"`
	AllowLateCheck       bool          `json:"allow_late_check" gorm:"default:true"`
	LateThresholdMinutes int           `json:"late_threshold_minutes" gorm:"default:15"`
	Location             string        `json:"location"`
	Notes                string        `json:"notes"`
	CreatedBy            uuid.UUID     `json:"created_by" gorm:"not null"`

	// Relations
	Classroom *Classroom `json:"classroom,omitempty" gorm:"foreignKey:ClassroomID"`
	Creator   *User      `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

// AttendanceRecord represents an individual attendance record
type AttendanceRecord struct {
	BaseModel
	SessionID       uuid.UUID        `json:"session_id" gorm:"not null"`
	StudentID       uuid.UUID        `json:"student_id" gorm:"not null"`
	Status          AttendanceStatus `json:"status" gorm:"type:attendance_status;not null"`
	CheckInTime     *time.Time       `json:"check_in_time"`
	CheckInMethod   *CheckInMethod   `json:"check_in_method" gorm:"type:check_in_method"`
	CheckInLocation string           `json:"check_in_location"`
	LateMinutes     int              `json:"late_minutes" gorm:"default:0"`
	Notes           string           `json:"notes"`
	MarkedBy        *uuid.UUID       `json:"marked_by"`
	IsModified      bool             `json:"is_modified" gorm:"default:false"`
	ModifiedAt      *time.Time       `json:"modified_at"`
	ModifiedBy      *uuid.UUID       `json:"modified_by"`

	// Relations
	Session  *AttendanceSession `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Student  *User              `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Marker   *User              `json:"marker,omitempty" gorm:"foreignKey:MarkedBy"`
	Modifier *User              `json:"modifier,omitempty" gorm:"foreignKey:ModifiedBy"`
}

// Message represents a message in the system
type Message struct {
	BaseModel
	SenderID        uuid.UUID     `json:"sender_id" gorm:"not null"`
	RecipientID     *uuid.UUID    `json:"recipient_id"`
	ClassroomID     *uuid.UUID    `json:"classroom_id"`
	ParentMessageID *uuid.UUID    `json:"parent_message_id"`
	Subject         string        `json:"subject"`
	Content         string        `json:"content" gorm:"not null"`
	MessageType     MessageType   `json:"message_type" gorm:"type:message_type;default:private"`
	Priority        PriorityLevel `json:"priority" gorm:"type:priority_level;default:normal"`
	IsRead          bool          `json:"is_read" gorm:"default:false"`
	ReadAt          *time.Time    `json:"read_at"`
	IsDeleted       bool          `json:"is_deleted" gorm:"default:false"`
	DeletedAt       *time.Time    `json:"deleted_at"`

	// Relations
	Sender        *User      `json:"sender,omitempty" gorm:"foreignKey:SenderID"`
	Recipient     *User      `json:"recipient,omitempty" gorm:"foreignKey:RecipientID"`
	Classroom     *Classroom `json:"classroom,omitempty" gorm:"foreignKey:ClassroomID"`
	ParentMessage *Message   `json:"parent_message,omitempty" gorm:"foreignKey:ParentMessageID"`
}

// Assignment represents an assignment
type Assignment struct {
	BaseModel
	ClassroomID         uuid.UUID        `json:"classroom_id" gorm:"not null"`
	Title               string           `json:"title" gorm:"not null"`
	Description         string           `json:"description"`
	Instructions        string           `json:"instructions"`
	AssignmentType      AssignmentType   `json:"assignment_type" gorm:"type:assignment_type;default:homework"`
	DueDate             *time.Time       `json:"due_date"`
	MaxScore            float64          `json:"max_score" gorm:"type:numeric(8,2);default:100.00"`
	Weight              float64          `json:"weight" gorm:"type:numeric(5,2);default:1.00"`
	AllowLateSubmission bool             `json:"allow_late_submission" gorm:"default:false"`
	LatePenaltyPercent  float64          `json:"late_penalty_percent" gorm:"type:numeric(5,2);default:0"`
	SubmissionFormat    SubmissionFormat `json:"submission_format" gorm:"type:submission_format;default:both"`
	MaxFileSizeMB       int              `json:"max_file_size_mb" gorm:"default:10"`
	AllowedFileTypes    string           `json:"allowed_file_types" gorm:"type:jsonb"`
	IsPublished         bool             `json:"is_published" gorm:"default:false"`
	Status              AssignmentStatus `json:"status" gorm:"type:assignment_status;default:draft"`
	CreatedBy           uuid.UUID        `json:"created_by" gorm:"not null"`

	// Relations
	Classroom *Classroom `json:"classroom,omitempty" gorm:"foreignKey:ClassroomID"`
	Creator   *User      `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

// AssignmentSubmission represents a student's submission for an assignment
type AssignmentSubmission struct {
	BaseModel
	AssignmentID   uuid.UUID        `json:"assignment_id" gorm:"not null"`
	StudentID      uuid.UUID        `json:"student_id" gorm:"not null"`
	SubmissionText string           `json:"submission_text"`
	Status         SubmissionStatus `json:"status" gorm:"type:submission_status;default:draft"`
	SubmittedAt    *time.Time       `json:"submitted_at"`
	IsLate         bool             `json:"is_late" gorm:"default:false"`
	LateMinutes    int              `json:"late_minutes" gorm:"default:0"`
	Score          *float64         `json:"score" gorm:"type:numeric(8,2)"`
	Feedback       string           `json:"feedback"`
	GradedBy       *uuid.UUID       `json:"graded_by"`
	GradedAt       *time.Time       `json:"graded_at"`

	// Relations
	Assignment *Assignment `json:"assignment,omitempty" gorm:"foreignKey:AssignmentID"`
	Student    *User       `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Grader     *User       `json:"grader,omitempty" gorm:"foreignKey:GradedBy"`
}

// FileUpload represents a file uploaded to the system
type FileUpload struct {
	BaseModel
	UploaderID    uuid.UUID    `json:"uploader_id" gorm:"not null"`
	FileName      string       `json:"file_name" gorm:"not null"`
	OriginalName  string       `json:"original_name" gorm:"not null"`
	FilePath      string       `json:"file_path" gorm:"not null"`
	FileSize      int          `json:"file_size" gorm:"not null"`
	FileType      string       `json:"file_type" gorm:"not null"`
	MimeType      string       `json:"mime_type" gorm:"not null"`
	FileCategory  FileCategory `json:"file_category" gorm:"type:file_category;default:other"`
	ReferenceType string       `json:"reference_type"`
	ReferenceID   *uuid.UUID   `json:"reference_id"`
	IsPublic      bool         `json:"is_public" gorm:"default:false"`
	Checksum      string       `json:"checksum"`

	// Relations
	Uploader *User `json:"uploader,omitempty" gorm:"foreignKey:UploaderID"`
}

// FilePermission represents permissions for file access
type FilePermission struct {
	BaseModel
	FileID         uuid.UUID      `json:"file_id" gorm:"not null"`
	UserID         *uuid.UUID     `json:"user_id"`
	ClassroomID    *uuid.UUID     `json:"classroom_id"`
	PermissionType PermissionType `json:"permission_type" gorm:"type:permission_type;not null"`
	GrantedBy      uuid.UUID      `json:"granted_by" gorm:"not null"`
	ExpiresAt      *time.Time     `json:"expires_at"`

	// Relations
	File      *FileUpload `json:"file,omitempty" gorm:"foreignKey:FileID"`
	User      *User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Classroom *Classroom  `json:"classroom,omitempty" gorm:"foreignKey:ClassroomID"`
	Granter   *User       `json:"granter,omitempty" gorm:"foreignKey:GrantedBy"`
}

// AcademicCalendar represents events in the academic calendar
type AcademicCalendar struct {
	BaseModel
	SchoolID          *uuid.UUID `json:"school_id"`
	Title             string     `json:"title" gorm:"not null"`
	Description       string     `json:"description"`
	EventType         EventType  `json:"event_type" gorm:"type:event_type;not null"`
	StartDate         time.Time  `json:"start_date" gorm:"type:date;not null"`
	EndDate           *time.Time `json:"end_date" gorm:"type:date"`
	IsRecurring       bool       `json:"is_recurring" gorm:"default:false"`
	RecurrencePattern string     `json:"recurrence_pattern" gorm:"type:jsonb"`
	AffectsAttendance bool       `json:"affects_attendance" gorm:"default:true"`
	CreatedBy         uuid.UUID  `json:"created_by" gorm:"not null"`

	// Relations
	School  *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	Creator *User   `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

// AttendanceAnalytics represents attendance analytics for a student
type AttendanceAnalytics struct {
	BaseModel
	ClassroomID        uuid.UUID `json:"classroom_id" gorm:"not null"`
	StudentID          uuid.UUID `json:"student_id" gorm:"not null"`
	MonthYear          string    `json:"month_year" gorm:"not null"`
	TotalSessions      int       `json:"total_sessions" gorm:"default:0"`
	PresentCount       int       `json:"present_count" gorm:"default:0"`
	AbsentCount        int       `json:"absent_count" gorm:"default:0"`
	LateCount          int       `json:"late_count" gorm:"default:0"`
	ExcusedCount       int       `json:"excused_count" gorm:"default:0"`
	AttendanceRate     float64   `json:"attendance_rate" gorm:"type:numeric(5,2);default:0"`
	AverageLateMinutes float64   `json:"average_late_minutes" gorm:"type:numeric(5,2);default:0"`

	// Relations
	Classroom *Classroom `json:"classroom,omitempty" gorm:"foreignKey:ClassroomID"`
	Student   *User      `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

// SystemSetting represents system-wide settings
type SystemSetting struct {
	BaseModel
	Key         string          `json:"key" gorm:"unique;not null"`
	Value       string          `json:"value"`
	DataType    DataTypeSetting `json:"data_type" gorm:"type:data_type_setting;default:string"`
	Description string          `json:"description"`
	IsEditable  bool            `json:"is_editable" gorm:"default:true"`
}

// SessionToken represents user session tokens
type SessionToken struct {
	BaseModel
	UserID           uuid.UUID `json:"user_id" gorm:"not null"`
	TokenHash        string    `json:"token_hash" gorm:"not null"`
	RefreshTokenHash string    `json:"refresh_token_hash"`
	ExpiresAt        time.Time `json:"expires_at" gorm:"not null"`
	IsRevoked        bool      `json:"is_revoked" gorm:"default:false"`
	IPAddress        string    `json:"ip_address" gorm:"type:inet"`
	UserAgent        string    `json:"user_agent"`

	// Relations
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// Notification represents a notification
type Notification struct {
	BaseModel
	RecipientID    uuid.UUID        `json:"recipient_id" gorm:"not null"`
	SenderID       *uuid.UUID       `json:"sender_id"`
	Title          string           `json:"title" gorm:"not null"`
	Message        string           `json:"message" gorm:"not null"`
	Type           NotificationType `json:"type" gorm:"type:notification_type;not null"`
	Priority       PriorityLevel    `json:"priority" gorm:"type:priority_level;default:normal"`
	ReferenceType  ReferenceType    `json:"reference_type" gorm:"type:reference_type"`
	ReferenceID    *uuid.UUID       `json:"reference_id"`
	IsRead         bool             `json:"is_read" gorm:"default:false"`
	ReadAt         *time.Time       `json:"read_at"`
	IsSent         bool             `json:"is_sent" gorm:"default:false"`
	SentAt         *time.Time       `json:"sent_at"`
	DeliveryMethod string           `json:"delivery_method" gorm:"type:jsonb"`

	// Relations
	Recipient *User `json:"recipient,omitempty" gorm:"foreignKey:RecipientID"`
	Sender    *User `json:"sender,omitempty" gorm:"foreignKey:SenderID"`
}
