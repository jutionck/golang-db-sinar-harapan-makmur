package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/utils/common"
	"net/http"
)

type VehicleController struct {
	router         *gin.Engine
	authMiddleware middleware.AuthTokenMiddleware
	usecase        usecase.VehicleUseCase
	api.BaseApi
}

func (v *VehicleController) createHandler(c *gin.Context) {
	var payload entity.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = uuid.New().String()
	if err := v.usecase.RegisterNewVehicle(payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	v.NewSuccessSingleResponseCreated(c, payload, "OK")
}

func (v *VehicleController) updateHandler(c *gin.Context) {
	var payload entity.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := v.usecase.UpdateVehicle(payload); err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	v.NewSuccessSingleResponse(c, payload, "OK")
}

func (v *VehicleController) listHandler(c *gin.Context) {
	requestQueryParams, err := common.ValidateRequestQueryParams(c)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	vehicles, paging, err := v.usecase.Paging(requestQueryParams)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var vehicleInterface []interface{}
	for _, v := range vehicles {
		vehicleInterface = append(vehicleInterface, v)
	}
	v.NewSuccessPageResponse(c, vehicleInterface, "OK", paging)
}

func (v *VehicleController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := v.usecase.GetVehicle(id)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	v.NewSuccessSingleResponse(c, vehicle, "OK")
}

func (v *VehicleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := v.usecase.DeleteVehicle(id)
	if err != nil {
		v.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewVehicleController(r *gin.Engine, usecase usecase.VehicleUseCase, authMiddleware middleware.AuthTokenMiddleware) *VehicleController {
	controller := VehicleController{
		router:         r,
		usecase:        usecase,
		authMiddleware: authMiddleware,
	}
	r.GET("/vehicles", controller.listHandler)
	r.GET("/vehicles/:id", controller.getByIDHandler)
	r.POST("/vehicles", authMiddleware.RequireToken(), controller.createHandler)
	r.PUT("/vehicles", authMiddleware.RequireToken(), controller.updateHandler)
	r.DELETE("/vehicles/:id", authMiddleware.RequireToken(), controller.deleteHandler)
	return &controller
}
