package main

import (
	"database/sql"
	"log"
	"net/http"

	config "github.com/ericoalmeida/go-orders_clean-arch/internal/configs"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/handlers"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/repositories"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/usecases"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfig()

	connStr := config.GetEnv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewPostgresOrderRepository(db)
	useCase := usecases.NewGetAllOrdersUsecase(repository)
	handler := handlers.NewGetAllOrdersHandler(useCase)

	http.HandleFunc("/orders", handler.ListOrders)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
