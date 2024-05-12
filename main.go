package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prashantpansare99/go_assignment/api"
	"github.com/prashantpansare99/go_assignment/db"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func main() {
	fmt.Println("In main")

	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

	// Initialize the database connection
	dbConnectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
    err := db.InitializeDB(dbConnectionString)
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
