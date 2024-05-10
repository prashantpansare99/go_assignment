package db

// InitialData represents the initial data for each company.
type InitialData struct {
    CompanyID      string
    FinancialsData interface{}
    SalesData      interface{}
    EmployeeStats  interface{}
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

// SetInitialData stores initial data for the given company in the database.
func SetInitialData(initialData *InitialData) error {
    _, err := db.Exec("INSERT INTO initial_data (company_id, financials_data, sales_data, employee_stats) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE financials_data = VALUES(financials_data), sales_data = VALUES(sales_data), employee_stats = VALUES(employee_stats)", initialData.CompanyID, initialData.FinancialsData, initialData.SalesData, initialData.EmployeeStats)
    return err
}