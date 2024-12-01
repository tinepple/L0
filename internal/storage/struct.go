package storage

import "time"

type Order struct {
	OrderUID          string    `db:"order_uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	Transaction       string    `db:"transaction"`
	RequestID         string    `db:"request_id"`
	Currency          string    `db:"currency"`
	Provider          string    `db:"provider"`
	Amount            int       `db:"amount"`
	PaymentDt         int       `db:"payment_dt"`
	Bank              string    `db:"bank"`
	DeliveryCost      int       `db:"delivery_cost"`
	GoodsTotal        int       `db:"goods_total"`
	CustomFee         int       `db:"custom_fee"`
	Id                string    `db:"id"`
	Name              string    `db:"name"`
	Phone             string    `db:"phone"`
	Zip               string    `db:"zip"`
	City              string    `db:"city"`
	Address           string    `db:"address"`
	Region            string    `db:"region"`
	Email             string    `db:"email"`
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	Shardkey          string    `db:"shard_key"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}

type Item struct {
	ChrtID      int    `db:"chrt_id"`
	TrackNumber string `db:"track_number"`
	Price       int    `db:"price"`
	Rid         string `db:"rid"`
	Name        string `db:"name"`
	Sale        int    `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  int    `db:"total_price"`
	NmID        int    `db:"nm_id"`
	Brand       string `db:"brand"`
	Status      int    `db:"status"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Message struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}
