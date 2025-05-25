package domain

type OrdersRepository interface {
	GetAll() ([]Order, error)
}
