package api

import (
	"encoding/json"
	"net/http"
	"github.com/prashantpansare99/go_assignment/db"
)

func SalesHandler(w http.ResponseWriter, r *http.Request) {
    companyID := r.URL.Query().Get("companyId")
	initialData, err := db.GetSalesData(companyID)
    if err != nil {
        return
    }
	json.NewEncoder(w).Encode(initialData)
}