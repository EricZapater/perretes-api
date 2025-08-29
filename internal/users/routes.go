package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *UserHandler) {
	roles := router.Group("/users")
	{		
		roles.PUT("/:id", handler.Update)
		roles.DELETE("/:id", handler.Delete)
		roles.POST("/change-password", handler.ChangePassword)				
	}
}

func RegisterPublicRoutes(router *gin.RouterGroup, handler *UserHandler) {
    router.POST("/register", handler.Create) // Ruta pública per crear usuaris
}