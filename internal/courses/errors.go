package courses

import "errors"

var (
	ErrCourseNotFound     = errors.New("course not found")
	ErrClassNotFound      = errors.New("class not found")
	ErrEnrollmentNotFound = errors.New("enrollment not found")
	ErrInvalidID          = errors.New("invalid ID")
	ErrInvalidRequest     = errors.New("invalid request")
)