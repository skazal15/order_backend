package service

import (
	"errors"
	"fmt"
	"order-service/internal/app/domain"
	"order-service/internal/app/dto"
	"order-service/internal/app/interfaces"
	"order-service/internal/app/order/models"
	"order-service/public/paging"
)

type OrderService struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderService(orderRepo interfaces.OrderRepository) *OrderService {
	return &OrderService{orderRepo}
}

func (or OrderService) OrderModelToDTO(md []models.Orders) (m []dto.Order) {
	mk := make([]dto.Order, 0)
	for _, data := range md {
		mk = append(mk, dto.Order{
			OrderName:       data.OrderName,
			CustomerCompany: data.CustomerCompanyName,
			CustomerName:    data.CustomerName,
			OrderDate:       data.OrderCreatedAt,
			DeliveredAmount: data.DeliveredQuantity,
			TotalAmount:     data.OrderQuantity,
		})
	}
	return mk
}

func (or OrderService) respPagination(resp *paging.Paginator) *domain.DataResponseWithPagination {
	var record = resp.Records.(*[]models.Orders)
	var data = or.OrderModelToDTO(*record)
	fmt.Println(data)
	var j interface{} = &data
	return &domain.DataResponseWithPagination{
		TotalRecord: resp.TotalRecord,
		TotalPage:   resp.TotalPage,
		Records:     &j,
		Offset:      resp.Offset,
		Limit:       resp.Limit,
		Page:        resp.Page,
		PrevPage:    resp.PrevPage,
		NextPage:    resp.NextPage,
	}
}

func (or OrderService) GetOrder(page int) (m *domain.DataResponseWithPagination, err error) {
	resp, errs := or.orderRepo.GetOrder(page)
	if errs != nil {
		return nil, errors.New("order not found")
	}
	return or.respPagination(resp), nil
}

func (or OrderService) GetOrderByOrderAndProduct(page int, Order string, Product string) (m *domain.DataResponseWithPagination, err error) {
	resp, err := or.orderRepo.GetOrderByOrderNameorProduct(page, Order, Product)
	if err != nil {
		return nil, errors.New("order not found")
	}

	return or.respPagination(resp), nil
}

func (or OrderService) GetOrderByDateRange(page int, start string, end string) (m *domain.DataResponseWithPagination, err error) {
	resp, err := or.orderRepo.GetOrderByDateRange(page, start, end)
	if err != nil {
		return nil, errors.New("order not found")
	}
	return or.respPagination(resp), nil
}
