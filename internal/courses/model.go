package courses

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	Classes     []Class   `json:"classes,omitempty"`
}

type Class struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CourseID    uuid.UUID `json:"course_id" db:"course_id"`
	VideoURL    string    `json:"video_url" db:"video_url"`
	MaterialURL string    `json:"material_url" db:"material_url"`
	Order       int       `json:"order" db:"order"`
	IsActive    bool      `json:"is_active" db:"is_active"`
}

type UserCourse struct {	
	EnrollmentID uuid.UUID `json:"enrollment_id" db:"enrollment_id"`
	UserID   uuid.UUID `json:"user_id" db:"user_id"`
	CourseID uuid.UUID `json:"course_id" db:"course_id"`
	Title	string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageURL	string    `json:"image_url" db:"image_url"`
	StartDate *time.Time `json:"start_date" db:"start_date"`
	Classes  []UserClassProgress `json:"classes,omitempty"`
}

type UserClassProgress struct {
	EnrollmentID uuid.UUID `json:"enrollment_id" db:"enrollment_id"`
	ClassID      uuid.UUID `json:"class_id" db:"class_id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	IsDone       bool      `json:"is_done" db:"is_done"`
}