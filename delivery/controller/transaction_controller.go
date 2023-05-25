package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type TransactionController struct {
	router  *gin.Engine
	usecase usecase.TransactionUseCase
}

func (e *TransactionController) createHandler(c *gin.Context) {
	var payload dto.TransactionRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// sebaiknya ini di taruh di usecase
	payload.Id = uuid.New().String()
	newTransactionPayload := entity.Transaction{
		Id:       payload.Id,
		Vehicle:  entity.Vehicle{Id: payload.VehicleId},
		Customer: entity.Customer{Id: payload.CustomerId},
		Employee: entity.Employee{Id: sql.NullString{String: payload.EmployeeId}},
		Type:     payload.Type,
		Qty:      payload.Qty,
	}
	if err := e.usecase.RegisterNewTransaction(newTransactionPayload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (e *TransactionController) listHandler(c *gin.Context) {
	vehicles, err := e.usecase.FindAllTransaction()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve employee data"})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (e *TransactionController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := e.usecase.FindTransactionById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve employee data"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func NewTransactionController(r *gin.Engine, usecase usecase.TransactionUseCase) *TransactionController {
	controller := TransactionController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/transactions", controller.listHandler)
	r.GET("/transactions/:id", controller.getByIDHandler)
	r.POST("/transactions", controller.createHandler)
	return &controller
}
