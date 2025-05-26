package graph

import (
	"context"
	"time"

	"github.com/ericoalmeida/go-orders_clean-arch/graphql/generated"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/usecases"
)

type Resolver struct {
	OrderUseCase usecases.GetAllOrdersUsecase
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetAllOrders(ctx context.Context) ([]*generated.Order, error) {
	orders, err := r.OrderUseCase.ListAll()
	if err != nil {
		return nil, err
	}

	var result []*generated.Order
	for _, o := range orders {
		result = append(result, &generated.Order{
			ID:           o.ID,
			Item:         o.Item,
			Customer:     o.Customer,
			PurchaseDate: o.PurchaseDate.Format(time.RFC3339),
			Price:        int(o.Price),
		})
	}
	return result, nil
}
