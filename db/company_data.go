package db

import (
	"database/sql"
	"encoding/json"
)

// InitialData represents the initial data for each company.
type FinancialsData struct {
	Revenue  int `json:"revenue"`
	Expenses int `json:"expenses"`
}

type SalesData struct {
	TotalSales   int `json:"totalSales"`
	AveragePrice int `json:"averagePrice"`
}

type EmployeeStats struct {
	TotalEmployees int `json:"totalEmployees"`
	AverageSalary  int `json:"averageSalary"`
}

type InitialData struct {
	CompanyID      string         `json:"companyID"`
	FinancialsData FinancialsData `json:"financialsData"`
	SalesData      SalesData      `json:"salesData"`
	EmployeeStats  EmployeeStats  `json:"employeeStats"`
}

// GetInitialData retrieves initial data for the given company from the database.
func GetInitialData(companyID string) (*InitialData, error) {
    var initialData InitialData
    err := db.QueryRow("SELECT company_id, financials_data, sales_data, employee_stats FROM initial_data WHERE company_id = ?", companyID).Scan(&initialData.CompanyID, &initialData.FinancialsData, &initialData.SalesData, &initialData.EmployeeStats)
    if err != nil {
        return nil, err
    }
    return &initialData, nil
}

func SetInitialData(initialData *InitialData) error {
    // Convert structs to JSON strings
    financialsJSON, err := json.Marshal(initialData.FinancialsData)
    if err != nil {
        return err
    }

    salesJSON, err := json.Marshal(initialData.SalesData)
    if err != nil {
        return err
    }

    employeeJSON, err := json.Marshal(initialData.EmployeeStats)
    if err != nil {
        return err
    }

    // Execute SQL statement to insert data
    _, err = db.Exec("INSERT INTO initial_data (company_id, financials_data, sales_data, employee_stats) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE financials_data = VALUES(financials_data), sales_data = VALUES(sales_data), employee_stats = VALUES(employee_stats)", initialData.CompanyID, financialsJSON, salesJSON, employeeJSON)
    return err
}

func InsertInitialData(companyID string, financialsData, salesData, employeeStats interface{}) error {
    financialsJSON, err := json.Marshal(financialsData)
    if err != nil {
        return err
    }

    salesJSON, err := json.Marshal(salesData)
    if err != nil {
        return err
    }

    employeeJSON, err := json.Marshal(employeeStats)
    if err != nil {
        return err
    }
    _, err = db.Exec(`
        INSERT INTO initial_data (company_id, financials_data, sales_data, employee_stats)
        VALUES (?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
        financials_data = VALUES(financials_data),
        sales_data = VALUES(sales_data),
        employee_stats = VALUES(employee_stats)
    `, companyID, financialsJSON, salesJSON, employeeJSON)
    return err
}

func GetEmployeeData(companyID string) (*EmployeeStats, error) {
    var jsonStr string
    err := db.QueryRow("SELECT employee_stats FROM initial_data WHERE company_id = ?", companyID).Scan(&jsonStr)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // Return nil if no data found, indicating no error
        }
        return nil, err
    }

    var employeeData EmployeeStats
    err = json.Unmarshal([]byte(jsonStr), &employeeData)
    if err != nil {
        return nil, err
    }

    return &employeeData, nil
}

func GetFinancialsData(companyID string) (*FinancialsData, error) {
    var jsonStr string
    err := db.QueryRow("SELECT financials_data FROM initial_data WHERE company_id = ?", companyID).Scan(&jsonStr)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // Return nil if no data found, indicating no error
        }
        return nil, err
    }

    var financialsData FinancialsData
    err = json.Unmarshal([]byte(jsonStr), &financialsData)
    if err != nil {
        return nil, err
    }

    return &financialsData, nil
}

func GetSalesData(companyID string) (*SalesData, error) {
    var jsonStr string
    err := db.QueryRow("SELECT sales_data FROM initial_data WHERE company_id = ?", companyID).Scan(&jsonStr)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // Return nil if no data found, indicating no error
        }
        return nil, err
    }

    var salesData SalesData
    err = json.Unmarshal([]byte(jsonStr), &salesData)
    if err != nil {
        return nil, err
    }

    return &salesData, nil
}