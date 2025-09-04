package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service CourseService
}

func NewCourseHandler(service CourseService) *CourseHandler {
	return &CourseHandler{
		service: service,
	}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var request CourseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := h.service.CreateCourse(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var request CourseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := h.service.UpdateCourse(c.Request.Context(), id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteCourse(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	course, err := h.service.FindCourseByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.service.FindAllCourses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) CreateClass(c *gin.Context) {
	var request ClassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	class, err := h.service.CreateClass(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

func(h *CourseHandler) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var request ClassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	class, err := h.service.UpdateClass(c.Request.Context(), id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

func(h *CourseHandler) DeleteClass(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteClass(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func(h *CourseHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")
	class, err := h.service.FindClassByID(c.Request.Context(), id)	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

func(h *CourseHandler) GetClassesByCourseID(c *gin.Context) {
	courseID := c.Param("course_id")
	classes, err := h.service.FindClassesByCourseID(c.Request.Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func(h *CourseHandler) GetCoursesByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	courses, err := h.service.FindCoursesByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) EnrollUserToCourse(c *gin.Context){
	var request EnrollmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enrollment, err := h.service.EnrollUserToCourse(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, enrollment)
}

func (h *CourseHandler) MarkClassAsDone(c *gin.Context){
	enrollmentID := c.Param("enrollment_id")
	classID := c.Param("class_id")	
	err := h.service.MarkClassAsDone(c.Request.Context(), enrollmentID, classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *CourseHandler) UnEnrollUserFromCourse(c *gin.Context){
	id := c.Param("enrollment_id")
	err := h.service.UnEnrollUserFromCourse(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
