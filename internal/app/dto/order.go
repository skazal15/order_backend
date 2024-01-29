package dto

import "time"

type Order struct {
	OrderName       string    `json:"order_name"`
	CustomerCompany string    `json:"customer_company"`
	CustomerName    string    `json:"customer_name"`
	OrderDate       time.Time `json:"order_date"`
	DeliveredAmount float64   `json:"delivered_amount"`
	TotalAmount     float64   `json:"total_amount"`
}
