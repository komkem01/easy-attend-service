package migrations

import (
	"github.com/komkem01/easy-attend-service/model"
)

func Models() []any {
	return []any{
		// Core entities
		(*model.Schools)(nil),
		(*model.Genders)(nil),
		(*model.Prefixes)(nil),
		(*model.Users)(nil),
		(*model.UserProfiles)(nil),
		(*model.Classrooms)(nil),
		(*model.ClassroomStudents)(nil),
		(*model.ClassroomMembers)(nil),

		// Attendance system
		(*model.AttendanceSessions)(nil),
		(*model.AttendanceRecords)(nil),
		(*model.AttendanceAnalytics)(nil),
		(*model.AttendanceRecordsArchive)(nil),
		(*model.AttendanceSessionsArchive)(nil),

		// Class management
		(*model.ClassSchedules)(nil),
		(*model.Assignments)(nil),
		(*model.AssignmentFiles)(nil),
		(*model.AssignmentSubmissions)(nil),

		// Communication
		(*model.Messages)(nil),
		(*model.Notifications)(nil),

		// Calendar and Events
		(*model.AcademicCalendar)(nil),

		// File management
		(*model.FileUploads)(nil),

		// Security and Authentication
		(*model.SessionTokens)(nil),
		(*model.UserSessions)(nil),
		(*model.UserRolePermissions)(nil),
		(*model.SecurityEvents)(nil),
		(*model.AuditLogs)(nil),

		// API management
		(*model.ApiKeys)(nil),
		(*model.ApiRateLimits)(nil),

		// System
		(*model.SystemSettings)(nil),
		(*model.MetricsData)(nil),
		(*model.SearchLogs)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
		// User defined enums
		`CREATE TYPE user_role AS ENUM ('student', 'teacher', 'admin', 'super_admin');`,
		`CREATE TYPE user_status AS ENUM ('active', 'inactive', 'suspended', 'pending');`,
		`CREATE TYPE gender AS ENUM ('male', 'female', 'other', 'prefer_not_to_say');`,
		`CREATE TYPE attendance_status AS ENUM ('present', 'absent', 'late', 'excused');`,
		`CREATE TYPE session_status AS ENUM ('scheduled', 'active', 'completed', 'cancelled');`,
		`CREATE TYPE session_method AS ENUM ('code', 'qr', 'manual', 'location');`,
		`CREATE TYPE check_in_method AS ENUM ('code', 'qr', 'manual', 'location', 'auto');`,
		`CREATE TYPE assignment_type AS ENUM ('homework', 'quiz', 'exam', 'project', 'lab');`,
		`CREATE TYPE submission_format AS ENUM ('text', 'file', 'both');`,
		`CREATE TYPE assignment_status AS ENUM ('draft', 'published', 'archived');`,
		`CREATE TYPE submission_status AS ENUM ('draft', 'submitted', 'graded', 'returned');`,
		`CREATE TYPE message_type AS ENUM ('private', 'broadcast', 'announcement');`,
		`CREATE TYPE priority_level AS ENUM ('low', 'normal', 'high', 'urgent');`,
		`CREATE TYPE event_type AS ENUM ('session', 'assignment_due', 'exam', 'holiday', 'meeting', 'other');`,
		`CREATE TYPE platform AS ENUM ('web', 'mobile', 'api', 'webhook');`,
		`CREATE TYPE file_category AS ENUM ('profile_picture', 'assignment', 'document', 'image', 'video', 'audio', 'other');`,
		`CREATE TYPE permission_type AS ENUM ('read', 'write', 'delete', 'admin', 'manage_users', 'manage_classrooms', 'view_reports');`,
		`CREATE TYPE notification_type AS ENUM ('info', 'warning', 'error', 'success', 'reminder', 'announcement');`,
		`CREATE TYPE delivery_status AS ENUM ('pending', 'sent', 'delivered', 'failed', 'cancelled');`,
		`CREATE TYPE delivery_channel AS ENUM ('in_app', 'email', 'sms', 'push');`,
		`CREATE TYPE reference_type AS ENUM ('attendance_session', 'assignment', 'announcement', 'grade', 'user', 'classroom');`,
		`CREATE TYPE metric_type AS ENUM ('counter', 'gauge', 'histogram', 'summary');`,
		`CREATE TYPE search_type AS ENUM ('user', 'classroom', 'session', 'assignment', 'global');`,
		`CREATE TYPE risk_level AS ENUM ('low', 'medium', 'high', 'critical');`,
		`CREATE TYPE data_type_setting AS ENUM ('string', 'integer', 'float', 'boolean', 'json', 'text');`,
		`CREATE TYPE classroom_status AS ENUM ('active', 'inactive', 'archived');`,
		`CREATE TYPE classroom_role AS ENUM ('student', 'teacher', 'assistant', 'observer');`,
		`CREATE TYPE member_status AS ENUM ('active', 'inactive', 'pending', 'removed');`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{
		// Add indexes for better performance
		`CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);`,
		`CREATE INDEX IF NOT EXISTS idx_users_school_id ON users(school_id);`,
		`CREATE INDEX IF NOT EXISTS idx_users_prefix_id ON users(prefix_id);`,
		`CREATE INDEX IF NOT EXISTS idx_users_gender_id ON users(gender_id);`,
		`CREATE INDEX IF NOT EXISTS idx_classrooms_teacher_id ON classrooms(teacher_id);`,
		`CREATE INDEX IF NOT EXISTS idx_classrooms_school_id ON classrooms(school_id);`,
		`CREATE INDEX IF NOT EXISTS idx_classroom_students_classroom_id ON classroom_students(classroom_id);`,
		`CREATE INDEX IF NOT EXISTS idx_classroom_students_student_id ON classroom_students(student_id);`,
		`CREATE INDEX IF NOT EXISTS idx_attendance_sessions_classroom_id ON attendance_sessions(classroom_id);`,
		`CREATE INDEX IF NOT EXISTS idx_attendance_records_session_id ON attendance_records(session_id);`,
		`CREATE INDEX IF NOT EXISTS idx_attendance_records_student_id ON attendance_records(student_id);`,
		`CREATE INDEX IF NOT EXISTS idx_assignments_classroom_id ON assignments(classroom_id);`,
		`CREATE INDEX IF NOT EXISTS idx_messages_sender_id ON messages(sender_id);`,
		`CREATE INDEX IF NOT EXISTS idx_messages_recipient_id ON messages(recipient_id);`,
		`CREATE INDEX IF NOT EXISTS idx_messages_classroom_id ON messages(classroom_id);`,
	}
}
