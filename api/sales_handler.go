package api

import (
	"encoding/json"
	"net/http"
	"github.com/prashantpansare99/go_assignment/db"
)

// SalesHandler handles requests to the /api/sales endpoint.
func SalesHandler(w http.ResponseWriter, r *http.Request) {
    // Step 1: Parse request parameters
    // Extract company ID from the request parameters
    companyID := r.URL.Query().Get("companyId")
	initialData, err := db.GetSalesData(companyID)
    if err != nil {
        // Handle error
        // Respond with appropriate error message
        return
    }
	json.NewEncoder(w).Encode(initialData)
}