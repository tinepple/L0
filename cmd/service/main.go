package main

import (
	"WBL0/internal/handler"
	"WBL0/internal/services/kafka"
	"WBL0/internal/services/orders"
	"WBL0/internal/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
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

	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("ping", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}
	defer partConsumer.Close()

	apiHandler := handler.New(allOrders)
	kafkaService := kafka.New(partConsumer, storageRepo, apiHandler)

	producer, err := sarama.NewSyncProducer(
		[]string{fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, nil,
	)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	message := storage.Message{
		OrderUID:    "ebfe6316-e5f8-f465-38eb-f79b97b798c5",
		TrackNumber: "123",
		Entry:       "123",
		Delivery: storage.Delivery{
			Name:    "123",
			Phone:   "123",
			Zip:     "123",
			City:    "123",
			Address: "123",
			Region:  "123",
			Email:   "123",
		},
		Payment: storage.Payment{
			Transaction:  "ebfe6316-e5f8-f465-38eb-f79b97b798c8",
			RequestID:    "123",
			Currency:     "123",
			Provider:     "123",
			Amount:       123,
			PaymentDt:    123,
			Bank:         "123bank",
			DeliveryCost: 123,
			GoodsTotal:   123,
			CustomFee:    123,
		},
		Items: []storage.Item{
			{
				ChrtID:      12,
				TrackNumber: "12",
				Price:       12,
				Rid:         "12",
				Name:        "12",
				Sale:        12,
				Size:        "12",
				TotalPrice:  12,
				NmID:        12,
				Brand:       "12",
				Status:      12,
			},
		},
		Locale:            "123",
		InternalSignature: "123",
		CustomerID:        "123",
		DeliveryService:   "123",
		Shardkey:          "123",
		SmID:              123,
		DateCreated:       time.Time{},
		OofShard:          "123",
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "ping",
		Value: sarama.ByteEncoder(bytes),
	}

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				_, _, err = producer.SendMessage(msg)
				if err != nil {
					log.Fatalf("Failed to create storage: %v", err)
				}
			}
		}
	}()

	go func() {
		if err = kafkaService.Consume(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APISERVER_PORT")), apiHandler); err != nil {
		log.Fatal(err)
	}
}
