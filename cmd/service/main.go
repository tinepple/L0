package main

import (
	"WBL0/internal/handler"
	"WBL0/internal/messages"
	"WBL0/internal/services/cache"
	"WBL0/internal/services/kafka"
	"WBL0/internal/services/orders"
	"WBL0/internal/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
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

	internalCache := cache.NewCache()
	for _, order := range allOrders {
		_ = internalCache.Set(order.OrderUID, order)
	}

	apiHandler := handler.New(internalCache)
	if err != nil {
		log.Fatalf("Failed to create handler: %v", err)
	}

	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("orders", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}
	defer partConsumer.Close()

	kafkaService := kafka.New(partConsumer, storageRepo, internalCache)

	go func() {
		if err = kafkaService.Consume(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	producer, err := sarama.NewSyncProducer(
		[]string{fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, nil,
	)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				bytes, err := json.Marshal(generateTestMessage())
				if err != nil {
					log.Fatalf("Failed to create storage: %v", err)
				}

				msg := &sarama.ProducerMessage{
					Topic: "orders",
					Value: sarama.ByteEncoder(bytes),
				}

				_, _, err = producer.SendMessage(msg)
				if err != nil {
					log.Fatalf("Failed to create storage: %v", err)
				}
			}
		}
	}()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APISERVER_PORT")), apiHandler); err != nil {
		log.Fatal(err)
	}
}

func generateTestMessage() messages.Order {
	return messages.Order{
		OrderUID:    uuid.New().String(),
		TrackNumber: randStr(10),
		Entry:       randStr(10),
		Delivery: messages.Delivery{
			Name:    randStr(10),
			Phone:   randStr(10),
			Zip:     randStr(10),
			City:    randStr(10),
			Address: randStr(10),
			Region:  randStr(10),
			Email:   randStr(10),
		},
		Payment: messages.Payment{
			Transaction:  uuid.New().String(),
			RequestID:    randStr(10),
			Currency:     randStr(10),
			Provider:     randStr(10),
			Amount:       rand.Intn(1000),
			PaymentDt:    rand.Intn(1000),
			Bank:         randStr(10),
			DeliveryCost: rand.Intn(1000),
			GoodsTotal:   rand.Intn(1000),
			CustomFee:    rand.Intn(1000),
		},
		Items: []messages.Item{
			{
				ChrtID:      rand.Intn(1000),
				TrackNumber: randStr(10),
				Price:       rand.Intn(1000),
				Rid:         randStr(10),
				Name:        randStr(10),
				Sale:        rand.Intn(1000),
				Size:        randStr(10),
				TotalPrice:  rand.Intn(1000),
				NmID:        rand.Intn(1000),
				Brand:       randStr(10),
				Status:      rand.Intn(1000),
			},
		},
		Locale:            randStr(10),
		InternalSignature: randStr(10),
		CustomerID:        randStr(10),
		DeliveryService:   randStr(10),
		Shardkey:          randStr(10),
		SmID:              rand.Intn(1000),
		DateCreated:       time.Now(),
		OofShard:          randStr(10),
	}
}

var number = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = number[rand.Intn(len(number))]
	}
	return string(b)
}
