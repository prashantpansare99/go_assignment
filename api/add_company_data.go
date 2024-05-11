package api

import (
	"encoding/json"
	"net/http"

	"github.com/prashantpansare99/go_assignment/db"
)

func AddCompanyData(w http.ResponseWriter, r *http.Request) {
    // Step 1: Parse request parameters
    // Extract company ID from the request parameters
	var initialData db.InitialData
	if err := json.NewDecoder(r.Body).Decode(&initialData); err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

	if initialData.CompanyID == "" {
		http.Error(w, "Invalid company data", http.StatusBadRequest)
        return
	}

	if err := db.InsertInitialData(initialData.CompanyID, initialData.FinancialsData, initialData.SalesData, initialData.EmployeeStats); err != nil {
        http.Error(w, "Failed to insert &initialData data into database", http.StatusInternalServerError)
        return
    }

	response := map[string]string{"message": "Data added successfully"}
	json.NewEncoder(w).Encode(response)

}