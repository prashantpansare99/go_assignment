package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prashantpansare99/go_assignment/db"
	"github.com/prashantpansare99/go_assignment/api"
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

    // Start HTTP server
    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }

	// Set initial data for a company
	// companyID := "company456"
	// initialData := generateRandomInitialData()
	// err = db.SetInitialData(&initialData)
	// if err != nil {
	// 	log.Fatalf("Failed to set initial data for company %s: %v", companyID, err)
	// }
	// fmt.Printf("Initial data set successfully for company %s\n", companyID)

	// Retrieve initial data for the same company
	// retrievedData, err := db.GetInitialData(companyID)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve initial data for company %s: %v", companyID, err)
	// }
	// fmt.Printf("Initial data retrieved for company %s:\n", companyID)
	// fmt.Printf("FinancialsData: %v\n", retrievedData.FinancialsData)
	// fmt.Printf("SalesData: %v\n", retrievedData.SalesData)
	// fmt.Printf("EmployeeStats: %v\n", retrievedData.EmployeeStats)
}

// generateRandomInitialData generates random initial data for demonstration purposes.
func generateRandomInitialData() db.InitialData {
	rand.Seed(time.Now().UnixNano())

	return db.InitialData{
		CompanyID:      "company456",
		FinancialsData: rand.Intn(10000),       // Random financials data
		SalesData:      rand.Intn(1000),         // Random sales data
		EmployeeStats:  rand.Intn(500),          // Random employee stats
	}
}