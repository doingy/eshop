package ports

import "github.com/doingy/eshop/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
}
