package api

import (
	"encoding/json"
	"net/http"
	"github.com/prashantpansare99/go_assignment/db"
)

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
    companyID := r.URL.Query().Get("companyId")
	initialData, err := db.GetEmployeeData(companyID)
    if err != nil {
        return
    }
	json.NewEncoder(w).Encode(initialData)
}