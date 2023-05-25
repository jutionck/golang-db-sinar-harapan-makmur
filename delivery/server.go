package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/controller"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/manager"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	controller.NewVehicleController(s.engine, s.ucManager.VehicleUseCase())
	controller.NewCustomerController(s.engine, s.ucManager.CustomerUseCase())
	controller.NewEmployeeController(s.engine, s.ucManager.EmployeeUseCase())
	controller.NewTransactionController(s.engine, s.ucManager.TransactionUseCase())
}

func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	// infra manager
	infraManager, _ := manager.NewInfraManager(c)
	// repo manager
	repoManager := manager.NewRepositoryManager(infraManager)
	// use case manager
	useCaseManager := manager.NewUseCaseManager(repoManager)
	r := gin.Default()
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{
		ucManager: useCaseManager,
		engine:    r,
		host:      host,
		log:       infraManager.Log(),
	}
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}
