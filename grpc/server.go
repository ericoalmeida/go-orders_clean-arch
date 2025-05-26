package grpc

import (
	"context"
	"net"
	"time"

	orderpb "github.com/ericoalmeida/go-orders_clean-arch/grpc/order"

	"github.com/ericoalmeida/go-orders_clean-arch/internal/usecases"
	"google.golang.org/grpc"
)

type OrderGRPCServer struct {
	orderpb.OrderServiceServer
	usecase usecases.GetAllOrdersUsecase
}

func NewServer(input usecases.GetAllOrdersUsecase) *OrderGRPCServer {
	return &OrderGRPCServer{usecase: input}
}

func (s *OrderGRPCServer) GetAllOrders(ctx context.Context, req *orderpb.Empty) (*orderpb.GetAllOrdersResponse, error) {
	orders, err := s.usecase.ListAll()
	if err != nil {
		return nil, err
	}

	var grpcOrders []*orderpb.Order
	for _, o := range orders {
		grpcOrders = append(grpcOrders, &orderpb.Order{
			Id:           o.ID,
			Item:         o.Item,
			Customer:     o.Customer,
			PurchaseDate: o.PurchaseDate.Format(time.RFC3339),
			Price:        o.Price,
		})
	}

	return &orderpb.GetAllOrdersResponse{Orders: grpcOrders}, nil
}

func RunGRPCServer(usecase usecases.GetAllOrdersUsecase, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, NewServer(usecase))

	return grpcServer.Serve(lis)
}
