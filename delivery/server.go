package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/controller"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/delivery/middleware"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
	"github.com/sirupsen/logrus"
)

type Server struct {
	vehicleUC     usecase.VehicleUseCase
	customerUC    usecase.CustomerUseCase
	employeeUC    usecase.EmployeeUseCase
	transactionUC usecase.TransactionUseCase
	engine        *gin.Engine
	host          string
	log           *logrus.Logger
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	controller.NewVehicleController(s.engine, s.vehicleUC)
	controller.NewCustomerController(s.engine, s.customerUC)
	controller.NewEmployeeController(s.engine, s.employeeUC)
	controller.NewTransactionController(s.engine, s.transactionUC)
}

func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	db := dbConn.Conn()

	r := gin.Default()
	logger := logrus.New()
	vehicleRepo := repository.NewVehicleRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	vehicleUC := usecase.NewVehicleUseCase(vehicleRepo)
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	employeeUC := usecase.NewEmployeeUseCase(employeeRepo)
	transactionUC := usecase.NewTransactionUseCase(transactionRepo, vehicleUC, customerUC, employeeUC)
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{
		vehicleUC:     vehicleUC,
		customerUC:    customerUC,
		employeeUC:    employeeUC,
		transactionUC: transactionUC,
		engine:        r,
		host:          host,
		log:           logger,
	}
}
