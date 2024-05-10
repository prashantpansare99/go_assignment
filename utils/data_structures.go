// utils/data_structures.go

package utils

// APIRequest represents an API request.
type APIRequest struct {
    CompanyID string
    Endpoint  string
}

// InitialData represents the initial data for a company.
type InitialData struct {
    // Define fields for initial data (dummy data for example)
    FinancialsData interface{}
    SalesData      interface{}
    EmployeeStats  interface{}
}

// WaitingRequests represents a queue of waiting requests for a company.
type WaitingRequests struct {
    Requests []APIRequest
    // You can add more fields as needed, such as a mutex for concurrency safety
}