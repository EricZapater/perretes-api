CREATE TABLE courses (
    id uuid PRIMARY KEY NOT NULL,
    title varchar(250) NOT NULL,
    description text,
    image_url varchar(500),
    is_active bool DEFAULT true
);

CREATE INDEX idx_courses_title ON courses(title);
CREATE INDEX idx_courses_is_active ON courses(is_active);

CREATE TABLE classes (
    id uuid PRIMARY KEY NOT NULL,
    course_id uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title varchar(250) NOT NULL,
    content text,
    video_url varchar(500),
    material_url varchar(500),
    "order" int NOT NULL,
    is_active bool DEFAULT true
);

CREATE INDEX idx_classes_course_id ON classes(course_id);
CREATE INDEX idx_classes_order ON classes(course_id, "order");
CREATE INDEX idx_classes_is_active ON classes(is_active);

CREATE TABLE course_enrollments (
    id uuid PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    start_date date DEFAULT CURRENT_DATE,
    is_active bool DEFAULT true
);

CREATE UNIQUE INDEX idx_course_enrollments_user_course 
    ON course_enrollments(user_id, course_id);

CREATE INDEX idx_course_enrollments_user_id ON course_enrollments(user_id);
CREATE INDEX idx_course_enrollments_course_id ON course_enrollments(course_id);

CREATE TABLE class_progress (
    id uuid PRIMARY KEY NOT NULL,
    enrollment_id uuid NOT NULL REFERENCES course_enrollments(id) ON DELETE CASCADE,
    class_id uuid NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
    is_done bool DEFAULT false
);

CREATE UNIQUE INDEX idx_class_progress_enrollment_class
    ON class_progress(enrollment_id, class_id);

CREATE INDEX idx_class_progress_enrollment_id ON class_progress(enrollment_id);
CREATE INDEX idx_class_progress_class_id ON class_progress(class_id);
