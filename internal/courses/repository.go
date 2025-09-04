package courses

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, course Course) (Course, error)
	UpdateCourse(ctx context.Context, course Course) (Course, error)
	DeleteCourse(ctx context.Context, id uuid.UUID) error
	FindCourseById(ctx context.Context, id uuid.UUID) (Course, error)
	FindAllCourses(ctx context.Context) ([]Course, error)
	CreateClass(ctx context.Context, class Class) (Class, error)
	UpdateClass(ctx context.Context, class Class) (Class, error)
	DeleteClass(ctx context.Context, id uuid.UUID) error
	FindClassById(ctx context.Context, id uuid.UUID) (Class, error)
	FindClassesByCourseId(ctx context.Context, courseID uuid.UUID) ([]Class, error)
	FindCoursesByUserID(ctx context.Context, userID uuid.UUID) ([]UserCourse, error)
	EnrollUserToCourse(ctx context.Context, userID, courseID uuid.UUID) (UserCourse, error)
	MarkClassAsDone(ctx context.Context, enrollmentID, classID uuid.UUID) error
	UnEnrollUserFromCourse(ctx context.Context, enrollmentID uuid.UUID) error
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) CreateCourse(ctx context.Context, course Course) (Course, error) {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO courses(id, title, description, image_url, is_active)
		VALUES($1, $2, $3, $4, $5)`,
		course.ID, course.Title, course.Description, course.ImageURL, course.IsActive,
	)
	if err != nil {
		return Course{}, err
	}
	return course, err
}

func (r *courseRepository) UpdateCourse(ctx context.Context, course Course) (Course, error){
	_, err := r.db.ExecContext(ctx, `
		UPDATE courses
		set title = $1,
		description = $2,
		image_url = $3,
		is_active = $4
		WHERE id = $5`,
		course.Title, course.Description, course.ImageURL, course.IsActive, course.ID,
		)
	if err != nil {
		return Course{}, err
	}
	return course, err
}

func (r *courseRepository) DeleteCourse(ctx context.Context, id uuid.UUID) error {
    _, err := r.db.ExecContext(ctx, `
        DELETE FROM courses WHERE id = $1`, id,
    )
    return err
}

func (r *courseRepository) FindCourseById(ctx context.Context, id uuid.UUID) (Course, error) {
    var course Course
    row := r.db.QueryRowContext(ctx, `
        SELECT id, title, description, image_url, is_active FROM courses WHERE id = $1`, id)
    err := row.Scan(&course.ID, &course.Title, &course.Description, &course.ImageURL, &course.IsActive)
    if err != nil {
        return Course{}, err
    }
    return course, nil
}

func (r *courseRepository) FindAllCourses(ctx context.Context) ([]Course, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT id, title, description, image_url, is_active FROM courses`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var courses []Course
    for rows.Next() {
        var c Course
        err = rows.Scan(&c.ID, &c.Title, &c.Description, &c.ImageURL, &c.IsActive)
        if err != nil {
            return nil, err
        }
        courses = append(courses, c)
    }
    return courses, nil
}

func (r *courseRepository) CreateClass(ctx context.Context, class Class) (Class, error) {
    _, err := r.db.ExecContext(ctx, `
        INSERT INTO classes (id, course_id, title, content, video_url, material_url, "order", is_active) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
        class.ID, class.CourseID, class.Title, class.Content, class.VideoURL, class.MaterialURL, class.Order, class.IsActive)
    if err != nil {
        return Class{}, err
    }
    return class, nil
}

func (r *courseRepository) UpdateClass(ctx context.Context, class Class) (Class, error) {
    _, err := r.db.ExecContext(ctx, `
        UPDATE classes
        SET title=$1, content=$2, video_url=$3, material_url=$4, "order"=$5, is_active=$6
        WHERE id=$7`,
        class.Title, class.Content, class.VideoURL, class.MaterialURL, class.Order, class.IsActive, class.ID)
    if err != nil {
        return Class{}, err
    }
    return class, nil
}

func (r *courseRepository) DeleteClass(ctx context.Context, id uuid.UUID) error {
    _, err := r.db.ExecContext(ctx, `
        DELETE FROM classes WHERE id=$1`, id)
    return err
}

func (r *courseRepository) FindClassById(ctx context.Context, id uuid.UUID) (Class, error) {
    var class Class
    row := r.db.QueryRowContext(ctx, `
        SELECT id, course_id, title, content, video_url, material_url, "order", is_active FROM classes WHERE id=$1`, id)
    err := row.Scan(&class.ID, &class.CourseID, &class.Title, &class.Content, &class.VideoURL, &class.MaterialURL, &class.Order, &class.IsActive)
    if err != nil {
        return Class{}, err
    }
    return class, nil
}

func (r *courseRepository) FindClassesByCourseId(ctx context.Context, courseID uuid.UUID) ([]Class, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT id, course_id, title, content, video_url, material_url, "order", is_active
        FROM classes WHERE course_id = $1 ORDER BY "order"`, courseID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var classes []Class
    for rows.Next() {
        var c Class
        err = rows.Scan(&c.ID, &c.CourseID, &c.Title, &c.Content, &c.VideoURL, &c.MaterialURL, &c.Order, &c.IsActive)
        if err != nil {
            return nil, err
        }
        classes = append(classes, c)
    }
    return classes, nil
}

func (r *courseRepository) FindCoursesByUserID(ctx context.Context, userID uuid.UUID) ([]UserCourse, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT ce.id as enrollment_id, c.id as course_id, c.title, c.description, c.image_url, ce.user_id, ce.start_date
        FROM course_enrollments ce
        JOIN courses c ON ce.course_id = c.id
        WHERE ce.user_id = $1 AND ce.is_active = true`, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var userCourses []UserCourse
    for rows.Next() {
        var uc UserCourse
        err = rows.Scan(&uc.EnrollmentID, &uc.CourseID, &uc.Title, &uc.Description, &uc.ImageURL, &uc.UserID, &uc.StartDate)
        if err != nil {
            return nil, err
        }

        // Fetch classes progress for this enrollment
        classRows, err := r.db.QueryContext(ctx, `
            SELECT cp.enrollment_id, cl.id, cl.title, cl.content, cp.is_done
            FROM class_progress cp
            JOIN classes cl ON cp.class_id = cl.id
            WHERE cp.enrollment_id = $1`, uc.EnrollmentID)
        if err != nil {
            return nil, err
        }
        var classesProgress []UserClassProgress
        for classRows.Next() {
            var ucp UserClassProgress
            var content string
            err = classRows.Scan(&ucp.EnrollmentID, &ucp.ClassID, &ucp.Title, &content, &ucp.IsDone)
            if err != nil {
                classRows.Close()
                return nil, err
            }
            ucp.Description = content
            classesProgress = append(classesProgress, ucp)
        }
        classRows.Close()
        uc.Classes = classesProgress
        userCourses = append(userCourses, uc)
    }
    return userCourses, nil
}

func (r *courseRepository) EnrollUserToCourse(ctx context.Context, userID, courseID uuid.UUID) (UserCourse, error) {
    var userCourse UserCourse
    _, err := r.db.ExecContext(ctx, `
        INSERT INTO course_enrollments(id, user_id, course_id)
        VALUES ($1, $2, $3)`, uuid.New(), userID, courseID)
    if err != nil {
        return userCourse, err
    }
    query := `
        SELECT ce.id as enrollment_id, c.id as course_id, c.title, c.description, c.image_url, ce.user_id, ce.start_date
        FROM course_enrollments ce
        JOIN courses c ON ce.course_id = c.id
        WHERE ce.user_id = $1 AND ce.course_id = $2 AND ce.is_active = true LIMIT 1`
    row := r.db.QueryRowContext(ctx, query, userID, courseID)
    err = row.Scan(&userCourse.EnrollmentID, &userCourse.CourseID, &userCourse.Title, &userCourse.Description, &userCourse.ImageURL, &userCourse.UserID, &userCourse.StartDate)
    if err != nil {
        return userCourse, err
    }
    return userCourse, nil
}

func (r *courseRepository) MarkClassAsDone(ctx context.Context, enrollmentID, classID uuid.UUID) error {
    _, err := r.db.ExecContext(ctx, `
        INSERT INTO class_progress (id, enrollment_id, class_id, is_done) VALUES ($1, $2, $3, true)
        ON CONFLICT (enrollment_id, class_id) DO UPDATE SET is_done = true
    `, uuid.New(), enrollmentID, classID)
    return err
}

func (r *courseRepository) UnEnrollUserFromCourse(ctx context.Context, enrollmentID uuid.UUID) error {
    _, err := r.db.ExecContext(ctx, `DELETE FROM course_enrollments WHERE id = $1`, enrollmentID)
    return err
}
