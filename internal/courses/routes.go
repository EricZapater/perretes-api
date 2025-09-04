package courses

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *CourseHandler) {
	courses := router.Group("/courses")
	{
		// CRUD Cursos
		courses.POST("", handler.CreateCourse)
		courses.PUT("/:id", handler.UpdateCourse)
		courses.DELETE("/:id", handler.DeleteCourse)
		courses.GET("/:id", handler.GetCourseByID)
		courses.GET("", handler.GetAllCourses)

		// CRUD Classes
		courses.POST("/classes", handler.CreateClass)
		courses.PUT("/classes/:id", handler.UpdateClass)
		courses.DELETE("/classes/:id", handler.DeleteClass)
		courses.GET("/classes/:id", handler.GetClassByID)
		courses.GET("/classes/bycourse/:course_id", handler.GetClassesByCourseID)

		// Enrolaments i progrés
		courses.POST("/enroll", handler.EnrollUserToCourse)
		courses.DELETE("/enroll/:enrollment_id", handler.UnEnrollUserFromCourse)
		courses.POST("/enroll/:enrollment_id/classes/:class_id/done", handler.MarkClassAsDone)

		// Recuperar cursos per usuari amb progrés
		courses.GET("/user/:user_id", handler.GetCoursesByUserID)
	}
}
