package main

import (
	"WBL0/internal/handler"
	"WBL0/internal/orders"
	"WBL0/internal/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	storageRepo, err := storage.New(db)
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}

	ordersService := orders.New(storageRepo)

	allOrders, err := ordersService.GetOrders(context.Background())
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}

	apiHandler := handler.New(allOrders)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APISERVER_PORT")), apiHandler); err != nil {
		log.Fatal(err)
	}

	//var dest int64
	//
	//err = db.QueryRowContext(context.Background(), db.Rebind("select id from asd where id = 5")).Scan(
	//	&dest,
	//)
	//if err != nil {
	//	log.Fatalf("QueryRowContext to connect to database: %v", err)
	//}

	//fmt.Println("УРА!", dest)
}
