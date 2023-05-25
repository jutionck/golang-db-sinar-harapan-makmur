package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type CustomerController struct {
	router  *gin.Engine
	usecase usecase.CustomerUseCase
}

func (cc *CustomerController) createHandler(c *gin.Context) {
	var payload entity.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	payload.Id = uuid.New().String()
	if err := cc.usecase.RegisterNewCustomer(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (cc *CustomerController) updateHandler(c *gin.Context) {
	var payload entity.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.usecase.UpdateCustomer(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payload)
}

func (cc *CustomerController) listHandler(c *gin.Context) {
	vehicles, err := cc.usecase.FindAllCustomer()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve customer data"})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (cc *CustomerController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := cc.usecase.GetCustomer(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve customer data"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (cc *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := cc.usecase.DeleteCustomer(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewCustomerController(r *gin.Engine, usecase usecase.CustomerUseCase) *CustomerController {
	controller := CustomerController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/customers", controller.listHandler)
	r.GET("/customers/:id", controller.getByIDHandler)
	r.POST("/customers", controller.createHandler)
	r.PUT("/customers", controller.updateHandler)
	r.DELETE("/customers/:id", controller.deleteHandler)
	return &controller
}
