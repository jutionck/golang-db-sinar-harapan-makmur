package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/api"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type UserController struct {
	router  *gin.Engine
	usecase usecase.UserUseCase
	api.BaseApi
}

func (u *UserController) createHandler(c *gin.Context) {
	var payload entity.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		u.NewErrorErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = uuid.New().String()
	if err := u.usecase.RegisterNewUser(payload); err != nil {
		u.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userCredentialDto := dto.UserCredentialResponseDto{
		Id:       payload.Id,
		UserName: payload.UserName,
	}
	u.NewSuccessSingleResponseCreated(c, userCredentialDto, "OK")
}

func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.usecase.FindAllUser()
	if err != nil {
		u.NewErrorErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []dto.UserCredentialResponseDto
	for _, user := range users {
		response = append(response, dto.UserCredentialResponseDto{
			Id:       user.Id,
			UserName: user.UserName,
		})
	}
	c.JSON(http.StatusOK, response)
}

func NewUserController(r *gin.Engine, usecase usecase.UserUseCase) *UserController {
	controller := UserController{
		router:  r,
		usecase: usecase,
	}
	r.POST("/users", controller.createHandler)
	r.GET("/users", controller.listHandler)
	return &controller
}
