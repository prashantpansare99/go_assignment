package utils

import (
	"fmt"
	"sync"
	"github.com/prashantpansare99/go_assignment/db"
)

type Request struct {
    CompanyID string
    API       string
}

type Response struct {
    CompanyID string
    API       string
    Result    interface{}
}

type RequestManager struct {
    requests map[string]map[string]chan Request
    mutex    sync.Mutex
}

var RequestManagerInstance = NewRequestManager()

func NewRequestManager() *RequestManager {
    return &RequestManager{
        requests: make(map[string]map[string]chan Request),
    }
}

func (rm *RequestManager) HandleRequest(req Request, initialData interface{}) chan Response {
    rm.mutex.Lock()
    defer rm.mutex.Unlock()

    fmt.Printf("Handling request for company %s, API %s\n", req.CompanyID, req.API)

    if _, ok := rm.requests[req.CompanyID]; !ok {
        rm.requests[req.CompanyID] = make(map[string]chan Request)
    }

    // If there's already a request for this company and API, return the existing channel
    if existingChan, ok := rm.requests[req.CompanyID][req.API]; ok {
		fmt.Println("Duplicate request detected for", req.CompanyID, req.API)
		respChan := make(chan Response)
		go func() {
			resp := <-existingChan // Wait for calculation completion
			respChan <- Response{CompanyID: resp.CompanyID, API: resp.API}
		}()
		return respChan
	}

    reqChan := make(chan Request)
    rm.requests[req.CompanyID][req.API] = reqChan

    go func() {
		calculate(req.CompanyID, req.API, initialData)
	
		reqChan <- Request{CompanyID: req.CompanyID, API: req.API}
	
		delete(rm.requests[req.CompanyID], req.API)
	
		close(reqChan)
	}()

    respChan := make(chan Response)
    go func() {
        resp := <-reqChan
        respChan <- Response{CompanyID: req.CompanyID, API: req.API, Result: resp}
    }()
    return respChan
}

func calculate(companyID string, api string, initialData interface{}) interface{} {
    financialsData, ok := initialData.(*db.FinancialsData)
    
    if !ok {
        return fmt.Errorf("initialData is not in expected format")
    }

    revenue := financialsData.Revenue
    expenses := financialsData.Expenses
    fmt.Println("revenue and expenses", revenue, expenses)
    profitMargin := (revenue - expenses) * 100 / revenue

    profitMarginFloat := float64(profitMargin)
    fmt.Println("Profit margin: ", profitMarginFloat)

    return fmt.Sprintf("Profit Margin for %s %s is: %.2f%%", companyID, api, profitMarginFloat)
}