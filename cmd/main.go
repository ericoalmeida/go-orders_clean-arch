package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/ericoalmeida/go-orders_clean-arch/graphql"
	"github.com/ericoalmeida/go-orders_clean-arch/graphql/generated"
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
	graphQlPort := config.GetEnv("GRAPHQL_PORT")

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

		log.Println("Listening on :" + port)
		if err := http.ListenAndServe(":"+port, mux); err != nil {
			log.Fatalf("Failed on trying to start HTTP server: %v", err)
		}
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

		log.Println("[gRPC] Listening on :" + grpcPort)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed on trying to start gRPC server: %v", err)
		}
	}()

	go func() {
		server := handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{
					Resolvers: &graph.Resolver{
						OrderUseCase: *useCase,
					},
				},
			),
		)

		http.Handle("/graphql", server)
		http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

		log.Println("[GraphQL] Listening on :" + graphQlPort)
		if err := http.ListenAndServe(":"+graphQlPort, nil); err != nil {
			log.Fatalf("GraphQL server failed: %v", err)
		}
	}()

	select {}
}
