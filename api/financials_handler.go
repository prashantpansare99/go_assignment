// financials_handler.go

package api

import (
	"encoding/json"
	"net/http"
	"github.com/prashantpansare99/go_assignment/db"
    "github.com/prashantpansare99/go_assignment/utils"
)

func FinancialsHandler(w http.ResponseWriter, r *http.Request) {
    
    companyID := r.URL.Query().Get("companyId")

    initialData, err := db.GetFinancialsData(companyID)
    if err != nil {
        // Handle error
        http.Error(w, "Failed to fetch financial data", http.StatusInternalServerError)
        return
    }

    req := utils.Request{CompanyID: companyID, API: "financials"}

    respChan := utils.RequestManagerInstance.HandleRequest(req, initialData)

    resp := <-respChan

    if err, ok := resp.Result.(error); ok && err != nil {
        // Handle error
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data, err := json.Marshal(initialData)
    if err != nil {
        http.Error(w, "Failed to encode response data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}
