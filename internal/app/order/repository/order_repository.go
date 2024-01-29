package repository

import (
	"fmt"
	"order-service/internal/app/order/models"
	"order-service/public/paging"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NeworderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) GetOrder(page int) (*paging.Paginator, error) {
	var m []models.Orders
	param := paging.Param{
		DB: o.db.Table("orders o").Select(
			"o.order_name, " +
				"cc.company_name AS customer_company_name, " +
				"c.name AS customer_name, " +
				"o.created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Australia/Melbourne' AS order_created_at, " +
				"COALESCE(coalesce (oi.price_per_unit,0) * d.delivered_quantity , 0) AS delivered_quantity, " +
				"COALESCE(coalesce(oi.price_per_unit,0)  * oi.quantity , 0) AS order_quantity",
		).Joins(
			"JOIN customers c ON o.customer_id = c.user_id",
		).Joins(
			"JOIN customer_companies cc ON c.company_id = cc.company_id",
		).Joins(
			"LEFT JOIN order_items oi ON o.id = oi.order_id",
		).Joins(
			"LEFT JOIN deliveries d ON oi.id = d.order_item_id",
		),
		Page:    page,
		Limit:   5,
		ShowSQL: true,
	}

	paginator := paging.Paging(&param, &m)
	return paginator, nil
}

func (o *orderRepository) GetOrderByDateRange(page int, start string, end string) (*paging.Paginator, error) {
	var m []models.Orders

	const layout = "2006-01-02"
	startTime, err := time.Parse(layout, start)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}
	endTime, err := time.Parse(layout, end)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %v", err)
	}

	param := paging.Param{
		DB: o.db.Table("orders o").Select(
			"o.order_name, "+
				"cc.company_name AS customer_company_name, "+
				"c.name AS customer_name, "+
				"o.created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Australia/Melbourne' AS order_created_at, "+
				"COALESCE(d.delivered_quantity * oi.price_per_unit , null) AS delivered_quantity, "+
				"COALESCE(coalesce(oi.price_per_unit,0)  * oi.quantity , 0) AS order_quantity",
		).Joins(
			"JOIN customers c ON o.customer_id = c.user_id",
		).Joins(
			"JOIN customer_companies cc ON c.company_id = cc.company_id",
		).Joins(
			"LEFT JOIN order_items oi ON o.id = oi.order_id",
		).Joins(
			"LEFT JOIN deliveries d ON oi.id = d.order_item_id",
		).Where("o.created_at >= ? AND o.created_at <= ?", startTime, endTime),
		Page:    page,
		Limit:   5,
		ShowSQL: true,
	}

	paginator := paging.Paging(&param, &m)
	return paginator, nil
}

func (o *orderRepository) GetOrderByOrderNameorProduct(page int, order string, product string) (*paging.Paginator, error) {

	var m []models.Orders

	param := paging.Param{
		DB: o.db.Table("orders o").Select(
			"o.order_name, "+
				"cc.company_name AS customer_company_name, "+
				"c.name AS customer_name, "+
				"o.created_at AT TIME ZONE 'UTC' AT TIME ZONE 'Australia/Melbourne' AS order_created_at, "+
				"COALESCE(d.delivered_quantity * oi.price_per_unit , null) AS delivered_quantity, "+
				"COALESCE(coalesce(oi.price_per_unit,0)  * oi.quantity , 0) AS order_quantity",
		).Joins(
			"JOIN customers c ON o.customer_id = c.user_id",
		).Joins(
			"JOIN customer_companies cc ON c.company_id = cc.company_id",
		).Joins(
			"LEFT JOIN order_items oi ON o.id = oi.order_id",
		).Joins(
			"LEFT JOIN deliveries d ON oi.id = d.order_item_id",
		).Where("o.order_name = ? OR oi.product = ?", order, product),
		Page:    page,
		Limit:   5,
		ShowSQL: true,
	}

	paginator := paging.Paging(&param, &m)
	return paginator, nil
}
