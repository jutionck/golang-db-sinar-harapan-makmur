package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type EmployeeController struct {
	router  *gin.Engine
	usecase usecase.EmployeeUseCase
	api.BaseApi
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var payload dto.EmployeeDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		e.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = uuid.New().String()
	newEmployee := entity.Employee{
		Id:          sql.NullString{String: payload.Id},
		FirstName:   sql.NullString{String: payload.FirstName},
		LastName:    sql.NullString{String: payload.LastName},
		Address:     sql.NullString{String: payload.Address},
		PhoneNumber: sql.NullString{String: payload.PhoneNumber},
		Email:       sql.NullString{String: payload.Email},
		Bod:         sql.NullTime{Time: payload.Bod},
		Position:    sql.NullString{String: payload.Position},
		Salary:      sql.NullInt64{Int64: payload.Salary},
		Manager:     &entity.Employee{Id: sql.NullString{String: payload.ManagerId}},
	}
	if err := e.usecase.RegisterNewEmployee(newEmployee); err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	e.NewSuccessSingleResponseCreated(c, payload, "OK")
}

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var payload entity.Employee
	if err := c.ShouldBindJSON(&payload); err != nil {
		e.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := e.usecase.UpdateEmployee(payload); err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	e.NewSuccessSingleResponse(c, payload, "OK")
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	vehicles, err := e.usecase.FindAllEmployee()
	if err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var vehicleResponses []dto.EmployeeDto
	for _, vehicle := range vehicles {
		vehicleResponses = append(vehicleResponses, dto.EmployeeDto{
			Id:          vehicle.Id.String,
			FirstName:   vehicle.FirstName.String,
			LastName:    vehicle.LastName.String,
			Position:    vehicle.Position.String,
			Address:     vehicle.Address.String,
			Email:       vehicle.Email.String,
			PhoneNumber: vehicle.PhoneNumber.String,
			Bod:         vehicle.Bod.Time,
			Salary:      vehicle.Salary.Int64,
			ManagerId:   vehicle.Manager.Id.String,
		})
	}
	c.JSON(http.StatusOK, vehicleResponses)
}

func (e *EmployeeController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := e.usecase.GetEmployee(id)
	if err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	vehicleResponse := dto.EmployeeDto{
		Id:          vehicle.Id.String,
		FirstName:   vehicle.FirstName.String,
		LastName:    vehicle.LastName.String,
		Position:    vehicle.Position.String,
		Address:     vehicle.Address.String,
		Email:       vehicle.Email.String,
		PhoneNumber: vehicle.PhoneNumber.String,
		Bod:         vehicle.Bod.Time,
		Salary:      vehicle.Salary.Int64,
		ManagerId:   vehicle.Manager.Id.String,
	}
	e.NewSuccessSingleResponse(c, vehicleResponse, "OK")
}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := e.usecase.DeleteEmployee(id)
	if err != nil {
		e.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewEmployeeController(r *gin.Engine, usecase usecase.EmployeeUseCase) *EmployeeController {
	controller := EmployeeController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/employees", controller.listHandler)
	r.GET("/employees/:id", controller.getByIDHandler)
	r.POST("/employees", controller.createHandler)
	r.PUT("/employees", controller.updateHandler)
	r.DELETE("/employees/:id", controller.deleteHandler)
	return &controller
}
