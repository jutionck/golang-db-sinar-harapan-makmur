package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type TransactionController struct {
	router         *gin.Engine
	authMiddleware middleware.AuthTokenMiddleware
	usecase        usecase.TransactionUseCase
	api.BaseApi
}

func (e *TransactionController) createHandler(c *gin.Context) {
	var payload dto.TransactionRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	e.NewSuccessSingleResponseCreated(c, payload, "OK")
}

func (e *TransactionController) listHandler(c *gin.Context) {
	transactions, err := e.usecase.FindAllTransaction()
	if err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (e *TransactionController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	transaction, err := e.usecase.FindTransactionById(id)
	if err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	e.NewSuccessSingleResponse(c, transaction, "OK")
}

func NewTransactionController(r *gin.Engine, usecase usecase.TransactionUseCase, authMiddleware middleware.AuthTokenMiddleware) *TransactionController {
	controller := TransactionController{
		router:         r,
		usecase:        usecase,
		authMiddleware: authMiddleware,
	}
	r.GET("/transactions", authMiddleware.RequireToken(), controller.listHandler)
	r.GET("/transactions/:id", authMiddleware.RequireToken(), controller.getByIDHandler)
	r.POST("/transactions", authMiddleware.RequireToken(), controller.createHandler)
	return &controller
}
