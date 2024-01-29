package models

import "time"

type Orders struct {
	OrderName           string    `json:"order_name"`
	CustomerCompanyName string    `json:"customer_company_name"`
	CustomerName        string    `json:"customer_name"`
	OrderCreatedAt      time.Time `json:"order_created_at"`
	DeliveredQuantity   float64   `json:"delivered_quantity"`
	OrderQuantity       float64   `json:"order_quantity"`
}
