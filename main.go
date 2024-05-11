package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prashantpansare99/go_assignment/api"
	"github.com/prashantpansare99/go_assignment/db"
)

func main() {
	fmt.Println("In main")
	// Initialize the database connection
	err := db.InitializeDB("root:root@tcp(127.0.0.1:3306)/company_info")
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer db.CloseDB()

	http.HandleFunc("/api/financials", api.FinancialsHandler)
	http.HandleFunc("/api/employees", api.EmployeeHandler)
	http.HandleFunc("/api/sales", api.SalesHandler)
	http.HandleFunc("/api/initialdata", api.AddCompanyData)

    // Start HTTP server
    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }

	// Below is the code to add data manually from main function 

	// initialData := db.InitialData{
	// 	CompanyID: "company456",
	// 	FinancialsData: db.FinancialsData{
	// 		Revenue: 20000,
	// 		Expenses: 9000,
	// 	},
	// 	SalesData: db.SalesData{
	// 		TotalSales:   600,
	// 		AveragePrice: 60,
	// 	},
	// 	EmployeeStats: db.EmployeeStats{
	// 		TotalEmployees: 80,
	// 		AverageSalary:  80000,
	// 	},
	// }

	// // Insert initial data into the database
	// err = db.SetInitialData(&initialData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
