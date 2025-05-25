package usecases

import "github.com/ericoalmeida/go-orders_clean-arch/internal/domain"

type GetAllOrdersUsecase struct {
	repository domain.OrdersRepository
}

func NewGetAllOrdersUsecase(input domain.OrdersRepository) *GetAllOrdersUsecase {
	return &GetAllOrdersUsecase{
		repository: input,
	}
}

func (usecase *GetAllOrdersUsecase) ListAll() ([]domain.Order, error) {
	return usecase.repository.GetAll()
}
