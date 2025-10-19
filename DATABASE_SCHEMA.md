# Easy Attend Service - Complete Database Schema

## Overview
This document outlines the complete database schema for the Easy Attend Service, a comprehensive attendance management system built with Go, Gin, and PostgreSQL.

## Database Models Created

### Core Entities
- **Schools**: Educational institutions
- **Users**: System users (students, teachers, admins)
- **UserProfiles**: Extended user information and preferences
- **Classrooms**: Class/course entities
- **ClassroomStudents**: Student enrollment in classrooms
- **ClassroomMembers**: General classroom membership with roles

### Attendance System
- **AttendanceSessions**: Individual attendance sessions
- **AttendanceRecords**: Student attendance records
- **AttendanceAnalytics**: Attendance analytics and statistics
- **AttendanceRecordsArchive**: Archived attendance records
- **AttendanceSessionsArchive**: Archived attendance sessions

### Class Management
- **ClassSchedules**: Class scheduling information
- **Assignments**: Course assignments
- **AssignmentFiles**: Files attached to assignments
- **AssignmentSubmissions**: Student assignment submissions

### Communication
- **Messages**: Messaging system between users
- **Notifications**: System notifications and alerts

### Calendar and Events
- **AcademicCalendar**: Academic calendar events

### File Management
- **FileUploads**: File upload management system

### Security and Authentication
- **SessionTokens**: User session token management
- **UserSessions**: Active user sessions tracking
- **UserRolePermissions**: User permissions system
- **SecurityEvents**: Security event logging
- **AuditLogs**: System audit trail

### API Management
- **ApiKeys**: API key management
- **ApiRateLimits**: API rate limiting configuration

### System Management
- **SystemSettings**: System configuration settings
- **MetricsData**: System metrics and analytics
- **SearchLogs**: Search activity logging

## Enum Types

The system uses comprehensive enum types for data consistency:

### User Related
- `user_role`: student, teacher, admin, super_admin
- `user_status`: active, inactive, suspended, pending
- `gender`: male, female, other, prefer_not_to_say

### Attendance Related
- `attendance_status`: present, absent, late, excused
- `session_status`: scheduled, active, completed, cancelled
- `session_method`: code, qr, manual, location
- `check_in_method`: code, qr, manual, location, auto

### Classroom Related
- `classroom_status`: active, inactive, archived
- `classroom_role`: student, teacher, assistant, observer
- `member_status`: active, inactive, pending, removed

### Assignment Related
- `assignment_type`: homework, quiz, exam, project, lab
- `assignment_status`: draft, published, archived
- `submission_status`: draft, submitted, graded, returned
- `submission_format`: text, file, both

### Communication Related
- `message_type`: private, broadcast, announcement
- `notification_type`: info, warning, error, success, reminder, announcement
- `delivery_status`: pending, sent, delivered, failed, cancelled
- `delivery_channel`: in_app, email, sms, push
- `priority_level`: low, normal, high, urgent

### System Related
- `event_type`: session, assignment_due, exam, holiday, meeting, other
- `platform`: web, mobile, api, webhook
- `file_category`: profile_picture, assignment, document, image, video, audio, other
- `permission_type`: read, write, delete, admin, manage_users, manage_classrooms, view_reports
- `reference_type`: attendance_session, assignment, announcement, grade, user, classroom
- `metric_type`: counter, gauge, histogram, summary
- `search_type`: user, classroom, session, assignment, global
- `risk_level`: low, medium, high, critical
- `data_type_setting`: string, integer, float, boolean, json, text

## Key Features

### UUID Primary Keys
All tables use UUID as primary keys for better scalability and security.

### Soft Deletes
Most models support soft deletes with `deleted_at` fields.

### Timestamps
All models include `created_at` and `updated_at` timestamps.

### Relationships
Proper foreign key relationships are defined using Bun ORM tags.

### JSON Fields
Complex data structures are stored as JSONB for flexibility.

### Indexes
Performance indexes are automatically created for frequently queried fields.

## Migration System

The migration system includes:
- Automatic enum type creation
- Table creation with proper constraints
- Index creation for performance
- Proper dependency ordering

## API Endpoints

The system includes REST API endpoints with:
- Authentication middleware
- CORS support
- PATCH and PUT method support
- Health check endpoints
- Proper response formatting

## CLI Commands

The application includes CLI commands for:
- `migrate`: Run database migrations
- `serve`: Start the HTTP server
- `healthcheck`: Check system health
- `version`: Show application version

## Technology Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **ORM**: Bun
- **Database**: PostgreSQL
- **Authentication**: JWT with bcrypt
- **CLI**: Cobra
- **UUID**: Google UUID package

## Getting Started

1. Set up environment variables
2. Run migrations: `./easy-attend migrate`
3. Start server: `./easy-attend serve`
4. Access health check: `http://localhost:8080/health`

## API Documentation

The API follows RESTful conventions with endpoints organized under `/api/v1/` prefix. All protected endpoints require JWT authentication via Authorization header.

## Environment Configuration

The application reads configuration from environment variables:
- `DATABASE_URL`: PostgreSQL connection string
- `JWT_SECRET`: JWT signing secret
- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin mode (debug/release)

This completes the comprehensive database schema and backend structure for the Easy Attend Service.