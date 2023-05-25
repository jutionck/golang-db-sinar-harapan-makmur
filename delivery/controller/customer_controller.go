package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type CustomerController struct {
	router  *gin.Engine
	usecase usecase.CustomerUseCase
	api.BaseApi
}

func (cc *CustomerController) createHandler(c *gin.Context) {
	var payload entity.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		cc.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = uuid.New().String()
	if err := cc.usecase.RegisterNewCustomer(payload); err != nil {
		cc.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	cc.NewSuccessSingleResponseCreated(c, payload, "OK")
}

func (cc *CustomerController) updateHandler(c *gin.Context) {
	var payload entity.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		cc.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := cc.usecase.UpdateCustomer(payload); err != nil {
		cc.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	cc.NewSuccessSingleResponse(c, payload, "OK")
}

func (cc *CustomerController) listHandler(c *gin.Context) {
	customers, err := cc.usecase.FindAllCustomer()
	if err != nil {
		cc.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	customer, err := cc.usecase.GetCustomer(id)
	if err != nil {
		cc.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	cc.NewSuccessSingleResponse(c, customer, "OK")
}

func (cc *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := cc.usecase.DeleteCustomer(id)
	if err != nil {
		cc.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
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
