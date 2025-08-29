package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService CustomerService
}

func NewCustomerHandler(customerService CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context){
	var request CustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group, err := h.customerService.Create(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var request CustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := h.customerService.Update(c.Request.Context(), id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	err := h.customerService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func(h *CustomerHandler) GetCustomerByID(c *gin.Context){
	id := c.Param("id")
	customer, err := h.customerService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func(h *CustomerHandler)GetAllCustomers(c *gin.Context){
	customers, err := h.customerService.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}
func(h *CustomerHandler) GetCustomerByUserID(c *gin.Context){
	userID := c.Param("user_id")
	customer, err := h.customerService.FindCustomerByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}