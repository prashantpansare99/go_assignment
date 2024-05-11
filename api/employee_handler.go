package api

import (
	"encoding/json"
	"net/http"
	"github.com/prashantpansare99/go_assignment/db"
)

// EmployeeHandler handles requests to the /api/employees endpoint.
func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
    // Step 1: Parse request parameters
    // Extract company ID from the request parameters
    companyID := r.URL.Query().Get("companyId")
	initialData, err := db.GetEmployeeData(companyID)
    if err != nil {
        return
    }
	json.NewEncoder(w).Encode(initialData)
}