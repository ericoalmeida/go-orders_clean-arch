package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/ericoalmeida/go-orders_clean-arch/grpc"
	grpcord "github.com/ericoalmeida/go-orders_clean-arch/grpc/order"
	config "github.com/ericoalmeida/go-orders_clean-arch/internal/configs"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/handlers"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/repositories"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/usecases"
	_ "github.com/lib/pq"
	googlerpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig()

	connStr := config.GetEnv("DATABASE_URL")
	port := config.GetEnv("PORT")
	grpcPort := config.GetEnv("GRPC_PORT")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewPostgresOrderRepository(db)
	useCase := usecases.NewGetAllOrdersUsecase(repository)

	go func() {
		mux := http.NewServeMux()
		orderHandler := handlers.NewGetAllOrdersHandler(useCase)
		mux.HandleFunc("/orders", orderHandler.ListOrders)

		log.Println("Listening on :8080")
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := googlerpc.NewServer()
		orderGrpcServer := grpc.NewServer(*useCase)
		grpcord.RegisterOrderServiceServer(grpcServer, orderGrpcServer)

		reflection.Register(grpcServer)

		log.Println("Listening on :" + grpcPort)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed on trying to start gRPC server: %v", err)
		}
	}()

	select {}
}
