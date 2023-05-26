package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"net/http"
)

type AuthController struct {
	router  *gin.Engine
	usecase usecase.AuthenticationUseCase
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload entity.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Println(payload.UserName, payload.Password)
	token, err := a.usecase.Login(payload.UserName, payload.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code":  http.StatusCreated,
		"token": token,
	})
}

func NewAuthController(r *gin.Engine, usecase usecase.AuthenticationUseCase) *AuthController {
	controller := AuthController{
		router:  r,
		usecase: usecase,
	}
	r.POST("/login", controller.loginHandler)
	return &controller
}
