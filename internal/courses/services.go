package courses

import (
	"context"

	"github.com/google/uuid"
)

type CourseService interface {
	CreateCourse(ctx context.Context, course CourseRequest) (Course, error)
	UpdateCourse(ctx context.Context, id string, course CourseRequest) (Course, error)
	DeleteCourse(ctx context.Context, id string) error
	FindCourseByID(ctx context.Context, id string) (Course, error)
	FindAllCourses(ctx context.Context) ([]Course, error)
	CreateClass(ctx context.Context, class ClassRequest) (Class, error)
	UpdateClass(ctx context.Context, id string, class ClassRequest) (Class, error)
	DeleteClass(ctx context.Context, id string) error
	FindClassByID(ctx context.Context, id string) (Class, error)
	FindClassesByCourseID(ctx context.Context, courseID string) ([]Class, error)
	FindCoursesByUserID(ctx context.Context, userID string) ([]UserCourse, error)
	EnrollUserToCourse(ctx context.Context, enrollment EnrollmentRequest) (UserCourse, error)
	MarkClassAsDone(ctx context.Context, enrollmentID, classID string) error
	UnEnrollUserFromCourse(ctx context.Context, enrollmentID string) error	
}

type courseService struct {
	repo CourseRepository
}

func NewCourseService(repo CourseRepository) CourseService {
	return &courseService{
		repo: repo,
	}
}

func(s *courseService) CreateCourse(ctx context.Context, course CourseRequest) (Course, error) {
	if course.Title == "" || course.Description == "" || course.ImageURL == "" {
		return Course{}, ErrInvalidRequest
	}
	
	newCourse := Course{
		ID:          uuid.New(),
		Title:       course.Title,
		Description: course.Description,
		ImageURL:    course.ImageURL,
		IsActive:    course.IsActive,
	}
	createdCourse, err := s.repo.CreateCourse(ctx, newCourse)
	if err != nil {
		return Course{}, err
	}
	return createdCourse, nil
}

func(s *courseService) UpdateCourse(ctx context.Context, id string, course CourseRequest) (Course, error){
	if id == "" || course.Title == "" || course.Description == "" || course.ImageURL == "" {
		return Course{}, ErrInvalidRequest
	}
	courseID, err := uuid.Parse(id)
	if err != nil {
		return Course{}, ErrInvalidID
	}
	updatedCourse := Course{
		ID:          courseID,
		Title:       course.Title,
		Description: course.Description,
		ImageURL:    course.ImageURL,
		IsActive:    course.IsActive,
	}
	result, err := s.repo.UpdateCourse(ctx, updatedCourse)
	if err != nil {
		return Course{}, err
	}
	return result, nil
}

func(s *courseService) DeleteCourse(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidRequest
	}
	courseID, err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidID
	}
	err = s.repo.DeleteCourse(ctx, courseID)
	if err != nil {
		return err
	}
	return nil
}

func(s *courseService) FindCourseByID(ctx context.Context, id string) (Course, error) {
	if id == "" {
		return Course{}, ErrInvalidRequest
	}
	courseID, err := uuid.Parse(id)
	if err != nil {
		return Course{}, ErrInvalidID
	}
	course, err := s.repo.FindCourseById(ctx, courseID)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func(s *courseService) FindAllCourses(ctx context.Context) ([]Course, error) {
	courses, err := s.repo.FindAllCourses(ctx)
	if err != nil {
		return nil, err
	}	
	return courses, nil
}

func(s *courseService) CreateClass(ctx context.Context, class ClassRequest) (Class, error) {
	if class.Title == "" || class.Content == "" || class.CourseID == "" || class.VideoURL == "" || class.MaterialURL == "" || class.Order <= 0 {
		return Class{}, ErrInvalidRequest
	}
	courseID, err := uuid.Parse(class.CourseID)
	if err != nil {
		return Class{}, ErrInvalidID
	}
	newClass := Class{
		ID:          uuid.New(),
		Title:       class.Title,
		Content:     class.Content,
		CourseID:    courseID,
		VideoURL:    class.VideoURL,
		MaterialURL: class.MaterialURL,
		Order:       class.Order,
		IsActive:    class.IsActive,
	}
	createdClass, err := s.repo.CreateClass(ctx, newClass)
	if err != nil {
		return Class{}, err
	}
	return createdClass, nil
}
func(s *courseService) UpdateClass(ctx context.Context, id string, class ClassRequest) (Class, error){
	if id == "" || class.Title == "" || class.Content == "" || class.CourseID ==
	 "" || class.VideoURL == "" || class.MaterialURL == "" || class.Order <= 0 {
		return Class{}, ErrInvalidRequest
	}
	classID, err := uuid.Parse(id)
	if err != nil {
		return Class{}, ErrInvalidID
	}
	courseID, err := uuid.Parse(class.CourseID)
	if err != nil {
		return Class{}, ErrInvalidID
	}
	updatedClass := Class{
		ID:          classID,
		Title:       class.Title,
		Content:     class.Content,
		CourseID:    courseID,
		VideoURL:    class.VideoURL,
		MaterialURL: class.MaterialURL,
		Order:       class.Order,
		IsActive:    class.IsActive,
	}
	result, err := s.repo.UpdateClass(ctx, updatedClass)
	if err != nil {
		return Class{}, err
	}
	return result, nil
}

func(s *courseService) DeleteClass(ctx context.Context, id string) error{
	if id == "" {
		return ErrInvalidRequest
	}
	classID , err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidID
	}
	err = s.repo.DeleteClass(ctx, classID)
	if err != nil {
		return err
	}
	return nil
}

func(s *courseService) FindClassByID(ctx context.Context, id string) (Class, error){
	if id == "" {
		return Class{}, ErrInvalidRequest
	}
	classID , err := uuid.Parse(id)
	if err != nil {
		return Class{}, ErrInvalidID
	}
	class, err := s.repo.FindClassById(ctx, classID)
	if err != nil {
		return Class{}, err
	}
	return class, nil
}

func(s *courseService) FindClassesByCourseID(ctx context.Context, courseID string) ([]Class, error){
	if courseID == "" {
		return nil, ErrInvalidRequest
	}
	parsedCourseID , err := uuid.Parse(courseID)
	if err != nil {
		return nil, ErrInvalidID
	}
	classes, err := s.repo.FindClassesByCourseId(ctx, parsedCourseID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func(s *courseService) FindCoursesByUserID(ctx context.Context, userID string) ([]UserCourse, error){
	if userID == "" {
		return nil, ErrInvalidRequest
	}
	parsedUserID , err := uuid.Parse(userID)
	if err != nil {
		return nil, ErrInvalidID
	}
	courses, err := s.repo.FindCoursesByUserID(ctx, parsedUserID)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func(s *courseService) EnrollUserToCourse(ctx context.Context, enrollment EnrollmentRequest) (UserCourse, error){
	if enrollment.UserID == "" || enrollment.CourseID == "" {
		return UserCourse{}, ErrInvalidRequest
	}
	userID, err := uuid.Parse(enrollment.UserID)
	if err != nil {
		return UserCourse{}, ErrInvalidID
	}
	courseID, err := uuid.Parse(enrollment.CourseID)
	if err != nil {
		return UserCourse{}, ErrInvalidID
	}
	userCourse, err := s.repo.EnrollUserToCourse(ctx, userID, courseID)
	if err != nil {
		return UserCourse{}, err
	}
	return userCourse, nil
}

func(s *courseService) MarkClassAsDone(ctx context.Context, enrollmentID, classID string) error{
	if enrollmentID == "" || classID == "" {
		return ErrInvalidRequest
	
	}
	parsedEnrollmentID , err := uuid.Parse(enrollmentID)
	if err != nil {
		return ErrInvalidID
	}
	parsedClassID , err := uuid.Parse(classID)
	if err != nil {
		return ErrInvalidID
	}
	err = s.repo.MarkClassAsDone(ctx, parsedEnrollmentID, parsedClassID)
	if err != nil {
		return err
	}
	return nil
}

func(s *courseService) UnEnrollUserFromCourse(ctx context.Context, enrollmentID string) error{
	if enrollmentID == "" {
		return ErrInvalidRequest
	}
	parsedEnrollmentID , err := uuid.Parse(enrollmentID)
	if err != nil {
		return ErrInvalidID
	}
	err = s.repo.UnEnrollUserFromCourse(ctx, parsedEnrollmentID)
	if err != nil {
		return err
	}
	return nil
}
	
