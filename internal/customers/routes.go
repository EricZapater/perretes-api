package customers

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *CustomerHandler){
	router.POST("/customers", handler.CreateCustomer)
	router.PUT("/customers/:id", handler.UpdateCustomer)
	router.DELETE("/customers/:id", handler.DeleteCustomer)
	router.GET("/customers/:id", handler.GetCustomerByID)
	router.GET("/customers", handler.GetAllCustomers)
	router.GET("/customers/user/:user_id", handler.GetCustomerByUserID)
}