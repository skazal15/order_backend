# Order Service
Service for get order

## Dependencies
- postgres
- golang 1.17
- Python 3.9.4
- psycopg2 2.9.9
- pandas 1.5.3

## API
### Order
- get all orders
- search order by order number or product name
- search order by range date

## Installation
- edit variables of seeder.py with your local settings in database_seeder file
- populate your database by running seeder.py (python seeder.py)
- running golang backend with command "go run .\cmd\api\main.go" 

## Project Structure

```
├── cmd
|   ├── api
│       └── main.go
├── database_seeder
|   ├── seeder.py
├── internal
|   ├── app
|   |   ├── domain
|   |   |   └── domain.go
|   |   ├── dto
|   |   |   └── order.go
|   |   ├── interface
|   |   |   └── order.go
|   |   ├── datahase
|   |   |   └── pgsql.go
|   |   ├── order
|   |   |   ├── repository
|   |   |   |   └── order_repository.go
|   |   |   ├── handler
|   |   |   |   └── order_handler.go
|   |   |   ├── service
|   |   |   |   └── order_service.go
|   |   |   ├── models
|   |   |   |   └── order.go
|   |   |   ├── handler
|   |   |   |   └── handler.go
|   |   |   ├── service/usecase
|   |   |   |   └── service.go
├── public
|   ├── common
|   |   ├── utils
|   |   |   └── error_response.go
|   |   |   └── struct_validator.go
|   ├── config
|   |   ├── variable.go
|   ├── paging
|   |   ├── paging.go
```