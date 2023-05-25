package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type VehicleController struct {
	router  *gin.Engine
	usecase usecase.VehicleUseCase
}

func (v *VehicleController) createHandler(c *gin.Context) {
	var payload entity.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	payload.Id = uuid.New().String()
	if err := v.usecase.RegisterNewVehicle(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (v *VehicleController) updateHandler(c *gin.Context) {
	var payload entity.Vehicle
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := v.usecase.UpdateVehicle(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payload)
}

func (v *VehicleController) listHandler(c *gin.Context) {
	vehicles, err := v.usecase.FindAllVehicle()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve vehicle data"})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (v *VehicleController) getByIDHandler(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := v.usecase.GetVehicle(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve vehicle data"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (v *VehicleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := v.usecase.DeleteVehicle(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewVehicleController(r *gin.Engine, usecase usecase.VehicleUseCase) *VehicleController {
	controller := VehicleController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/vehicles", controller.listHandler)
	r.GET("/vehicles/:id", controller.getByIDHandler)
	r.POST("/vehicles", controller.createHandler)
	r.PUT("/vehicles", controller.updateHandler)
	r.DELETE("/vehicles/:id", controller.deleteHandler)
	return &controller
}
