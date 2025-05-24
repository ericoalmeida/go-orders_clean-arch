package usecases

import "github.com/ericoalmeida/go-orders_clean-arch/internal/domain"

type OrderRepository interface {
	GetAll() ([]domain.Order, error)
}

type OrderUsecase struct {
	repository OrderRepository
}

func NewOrderUsecase(repository OrderRepository) *OrderUsecase {
	return &OrderUsecase{
		repository: repository,
	}
}

func (usecase *OrderUsecase) ListAll() ([]domain.Order, error) {
	return usecase.repository.GetAll()
}
