package courses

type CourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"required"`
	IsActive    bool   `json:"is_active"`
}

type ClassRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CourseID    string `json:"course_id" binding:"required"`
	VideoURL    string `json:"video_url" binding:"required"`
	MaterialURL string `json:"material_url" binding:"required"`
	Order       int    `json:"order" binding:"required"`
	IsActive    bool   `json:"is_active"`
}

type EnrollmentRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	CourseID string `json:"course_id" binding:"required"`
}