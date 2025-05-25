package main

import (
	"database/sql"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	config "github.com/ericoalmeida/go-orders_clean-arch/internal/configs"
	"github.com/ericoalmeida/go-orders_clean-arch/internal/domain"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfig()

	strConn := config.GetEnv("DATABASE_URL")

	db, err := sql.Open("postgres", strConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var recordsCount int
	err = db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&recordsCount)
	if err != nil {
		log.Fatal(err)
	}

	if recordsCount > 0 {
		return
	}

	commandSql := `INSERT INTO orders (item, customer, purchaseDate, price) VALUES ($1, $2, $3, $4)`

	for i := 0; i <= 9; i++ {
		order := domain.Order{
			Item:         gofakeit.ProductName(),
			Customer:     gofakeit.Name(),
			PurchaseDate: gofakeit.Date(),
			Price:        int64(gofakeit.Product().Price),
		}

		_, err = db.Exec(commandSql, order.Item, order.Customer, order.PurchaseDate, order.Price)

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed completed.")
}
