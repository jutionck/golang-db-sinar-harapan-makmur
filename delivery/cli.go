package delivery

import (
	"database/sql"
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/model/entity"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/usecase"
)

func VehicleCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	vehicleRepository := repository.NewVehicleRepository(dbConn.Conn())
	vehicleUseCase := usecase.NewVehicleUseCase(vehicleRepository)

	// cretae
	newVehicle := entity.Vehicle{
		Brand:     "Toyota",
		Model:     "Alphard",
		Color:     "Putih",
		Stock:     0,
		Status:    "Baru",
		SalePrice: 900000000,
	}
	newVehicle.SetId()
	if err := vehicleUseCase.RegisterNewVehicle(newVehicle); err != nil {
		fmt.Println(err)
		return
	}

	// // Get All
	// vehicles, err := vehicleUseCase.FindAllVehicle()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, v := range vehicles {
	// 		fmt.Println("ID:", v.Id)
	// 		fmt.Println("Brand:", v.Brand)
	// 		fmt.Println("Model:", v.Model)
	// 		fmt.Println("Production Year:", v.ProductionYear)
	// 		fmt.Println("Color:", v.Color)
	// 		fmt.Println("Price:", v.SalePrice)
	// 		fmt.Println("Is Automatic:", v.IsAutomatic)
	// 		fmt.Println("Stock:", v.Stock)
	// 		fmt.Println("Status:", v.Status)
	// 		fmt.Println()
	// 	}
	// }

	// Paging
	// requestQueryParams := dto.RequestQueryParams{
	// 	QueryParams: dto.QueryParams{
	// 		Order: "model",
	// 		Sort:  "asc",
	// 	},
	// 	PaginationParam: dto.PaginationParam{
	// 		Page:   1,
	// 		Offset: 0,
	// 		Limit:  5,
	// 	},
	// }
	// vehicles, paging, err := vehicleUseCase.Paging(requestQueryParams)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, v := range vehicles {
	// 		fmt.Println("ID:", v.Id)
	// 		fmt.Println("Brand:", v.Brand)
	// 		fmt.Println("Model:", v.Model)
	// 		fmt.Println("Production Year:", v.ProductionYear)
	// 		fmt.Println("Color:", v.Color)
	// 		fmt.Println("Price:", v.SalePrice)
	// 		fmt.Println("Is Automatic:", v.IsAutomatic)
	// 		fmt.Println("Stock:", v.Stock)
	// 		fmt.Println("Status:", v.Status)
	// 		fmt.Println()
	// 	}
	// 	fmt.Println("Paging:")
	// 	fmt.Println("page:", paging.Page)
	// 	fmt.Println("totalPages:", paging.TotalPages)
	// 	fmt.Println("totalRows:", paging.TotalRows)
	// 	fmt.Println("rowsPerPage:", paging.RowsPerPage)
	// }

	// Group By
	whereBy := map[string]interface{}{"production_year": 2022}
	vehiclesGroupCount, err := vehicleUseCase.GroupBy("brand", whereBy, "brand")
	if err != nil {
		// handle error
	} else {
		for _, v := range vehiclesGroupCount {
			fmt.Printf("Brand: %s, Total: %d\n", v.FieldName, v.FieldCount)
		}
	}
}

func CustomerCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	customerRepository := repository.NewCustomerRepository(dbConn.Conn())
	customerUseCase := usecase.NewCustomerUseCase(customerRepository)

	// cretae
	// bod, _ := time.Parse("2006-01-02", "1999-11-11")
	// newCustomer := entity.Customer{
	// 	FirstName:   "Tika",
	// 	LastName:    "Yesi",
	// 	PhoneNumber: "0821444444",
	// 	Email:       "tika.yesi@gmail.com",
	// 	Bod:         bod,
	// }
	// newCustomer.SetId()
	// if err := customerUseCase.RegisterNewCustomer(newCustomer); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Get All
	// customers, err := customerUseCase.FindAllCustomer()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, c := range customers {
	// 		fmt.Println("ID:", c.Id)
	// 		fmt.Println("Name:", c.FirstName, c.LastName)
	// 		fmt.Println("Phone Number:", c.PhoneNumber)
	// 		fmt.Println("Email:", c.Email)
	// 		fmt.Println("Birth Date:", c.Bod)
	// 		fmt.Println()
	// 	}
	// }

	// Find By Email
	customer, err := customerUseCase.FindCustomerByEmail("igladdinh@archive.org")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ID:", customer.Id)
		fmt.Println("Email:", customer.Email)
		fmt.Println()
	}
}

func EmployeeCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	employeeRepository := repository.NewEmployeeRepository(dbConn.Conn())
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository)

	// manager, err := employeeUseCase.FindManagerById("34258ecc-b35c-4da9-8574-c452475af11f")
	// bod, _ := time.Parse("2006-01-02", "1990-11-11")
	// newEmployee := entity.Employee{
	// 	FirstName:   sql.NullString{String: "Doni"},
	// 	LastName:    sql.NullString{String: "Octa"},
	// 	PhoneNumber: sql.NullString{String: "082929193738"},
	// 	Email:       sql.NullString{String: "doni.octo@gmail.com"},
	// 	Bod:         sql.NullTime{Time: bod},
	// 	Position:    sql.NullString{String: "Software Developer"},
	// 	Salary:      sql.NullInt64{Int64: 15000000},
	// 	Manager:     &manager,
	// }
	// newEmployee.SetId()
	// if err := employeeUseCase.RegisterNewEmployee(newEmployee); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Get All
	employees, err := employeeUseCase.FindAllEmployee()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, e := range employees {
			fmt.Println("ID:", e.Id)
			fmt.Println("Name:", e.FirstName, e.LastName)
			fmt.Println("Phone Number:", e.PhoneNumber)
			fmt.Println("Email:", e.Email)
			fmt.Println("Birth Date:", e.Bod)
			fmt.Println("Position:", e.Position)
			fmt.Println("Salary:", e.Salary)
			fmt.Println("Manager:", e.Manager)
			fmt.Println()
		}
	}

	// Get
	// e, err := employeeUseCase.GetEmployee("111111")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("ID:", e.Id)
	// fmt.Println("Name:", e.FirstName, e.LastName)
	// fmt.Println("Phone Number:", e.PhoneNumber)
	// fmt.Println("Email:", e.Email)
	// fmt.Println("Birth Date:", e.Bod)
	// fmt.Println("Position:", e.Position)
	// fmt.Println("Salary:", e.Salary)
	// fmt.Println("Manager:", e.Manager)
	// fmt.Println()
}

func TransactionCLI() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbConn, _ := config.NewDbConnection(c)
	defer dbConn.Conn().Close()

	transactionRepo := repository.NewTransactionRepository(dbConn.Conn())
	vehicleUseCase := usecase.NewVehicleUseCase(repository.NewVehicleRepository(dbConn.Conn()))
	customerUseCase := usecase.NewCustomerUseCase(repository.NewCustomerRepository(dbConn.Conn()))
	employeeUseCase := usecase.NewEmployeeUseCase(repository.NewEmployeeRepository(dbConn.Conn()))

	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo, vehicleUseCase, customerUseCase, employeeUseCase)
	newTransaction := entity.Transaction{
		Vehicle:  entity.Vehicle{Id: "b3a41ff7-a5af-4f04-b0e5-19e7451a8556"},
		Customer: entity.Customer{Id: "afc49d21-a381-42f0-8f0b-d94d4148d8e1"},
		Employee: entity.Employee{Id: sql.NullString{String: "15c68c8f-eff0-42cc-a8dd-903be384fa8a"}},
		Type:     "Online",
		Qty:      1,
	}
	newTransaction.SetId()
	if err := transactionUseCase.RegisterNewTransaction(newTransaction); err != nil {
		fmt.Println(err)
		return
	}

	// FindAllTransaction
	// transactions, err := transactionUseCase.FindAllTransaction()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, t := range transactions {
	// 		fmt.Println("ID:", t.Id)
	// 		fmt.Println("Date:", t.TransactionDate)
	// 		fmt.Println("Customer:", t.CustomerName)
	// 		fmt.Println("Vehicle:", t.VehicleBrand, t.VehicleModel)
	// 		fmt.Println("Employee:", t.EmployeeName)
	// 		fmt.Println("Type:", t.Type)
	// 		fmt.Println("Payment Amount:", t.PaymentAmount)
	// 		fmt.Println("Qty:", t.Qty)
	// 		fmt.Println()
	// 	}
	// }

	// Get
	// t, err := transactionUseCase.FindTransactionById("7c95d626-16a4-41fb-8ded-a54cf4e58f94")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("ID:", t.Id)
	// 	fmt.Println("Date:", t.TransactionDate)
	// 	fmt.Println("Customer:", t.CustomerName)
	// 	fmt.Println("Vehicle:", t.VehicleBrand, t.VehicleModel)
	// 	fmt.Println("Employee:", t.EmployeeName)
	// 	fmt.Println("Type:", t.Type)
	// 	fmt.Println("Payment Amount:", t.PaymentAmount)
	// 	fmt.Println("Qty:", t.Qty)
	// 	fmt.Println()
	// }

	// FindAllTransaction
	// transactions, err := transactionUseCase.GetAll()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, t := range transactions {
	// 		fmt.Println("ID:", t.Id)
	// 		fmt.Println("Date:", t.TransactionDate)
	// 		fmt.Println("Customer:", t.Customer)
	// 		fmt.Println("Vehicle:", t.Vehicle)
	// 		fmt.Println("Employee:", t.Employee)
	// 		fmt.Println("Employee Manager:", t.Employee.Manager)
	// 		fmt.Println("Type:", t.Type)
	// 		fmt.Println("Payment Amount:", t.PaymentAmount)
	// 		fmt.Println("Qty:", t.Qty)
	// 		fmt.Println()
	// 	}
	// }
}
