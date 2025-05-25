package repositories

import (
	"database/sql"

	"github.com/ericoalmeida/go-orders_clean-arch/internal/domain"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(input *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{
		db: input,
	}
}

func (repository *PostgresOrderRepository) GetAll() ([]domain.Order, error) {
	rows, err := repository.db.Query("SELECT ord.id, ord.item, ord.customer, ord.purchaseDate, ord.price FROM orders ord")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []domain.Order{}
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.Item, &order.Customer, &order.PurchaseDate, &order.Price); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
