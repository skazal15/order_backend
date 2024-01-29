package interfaces

import (
	"order-service/internal/app/domain"
	"order-service/public/paging"
)

type OrderRepository interface {
	GetOrder(page int) (*paging.Paginator, error)
	GetOrderByOrderNameorProduct(page int, order string, product string) (*paging.Paginator, error)
	GetOrderByDateRange(page int, start string, end string) (*paging.Paginator, error)
}

type OrderService interface {
	GetOrder(page int) (m *domain.DataResponseWithPagination, err error)
	GetOrderByOrderAndProduct(page int, Order string, Product string) (m *domain.DataResponseWithPagination, err error)
	GetOrderByDateRange(page int, start string, end string) (m *domain.DataResponseWithPagination, err error)
}
