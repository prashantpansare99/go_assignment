git clone https://github.com/prashantpansare99/go_assignment.git
cd go_assignment/
go mod tidy
go run ./main.go

Added postman collection json file in root structure
Import postman collection and run API's

.env file need below parameters
DB_USER=
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=company_info

table name is initial_data
have added sql queries manually

Endpoint 1
URL: /api/financials
Method: GET
Description: This endpoint will fetch financials data by companyId from MySQL

Endpoint 2
URL: /api/employees
Method: GET
Description: This endpoint will fetch employees data by companyId from MySQL

Endpoint 3
URL: /api/sales
Method: GET
Description: This endpoint will fetch sales data by companyId from MySQL

Endpoint 4
URL: /api/initialdata
Method: POST
Description: This endpoint will add comapany data in MySQL

api/
this is for handling api's

db/
this will handling connection with db

utils/
for concurrency handlers, data structures, error handlers