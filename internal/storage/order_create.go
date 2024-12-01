package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateOrder(ctx context.Context, model Message) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query, params, err := sq.Insert("deliveries").
		Columns(
			"name",
			"phone",
			"zip",
			"city",
			"address",
			"region",
			"email",
		).
		Values(
			model.Delivery.Name,
			model.Delivery.Phone,
			model.Delivery.Zip,
			model.Delivery.City,
			model.Delivery.Address,
			model.Delivery.Region,
			model.Delivery.Email,
		).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	var deliveryId string

	err = tx.QueryRowContext(ctx, tx.Rebind(query), params...).Scan(
		&deliveryId,
	)
	if err != nil {
		return err
	}

	query, params, err = sq.Insert("payments").
		Columns(
			"transaction",
			"request_id",
			"currency",
			"provider",
			"amount",
			"payment_dt",
			"bank",
			"delivery_cost",
			"goods_total",
			"custom_fee",
		).
		Values(
			model.Payment.Transaction,
			model.Payment.RequestID,
			model.Payment.Currency,
			model.Payment.Provider,
			model.Payment.Amount,
			model.Payment.PaymentDt,
			model.Payment.Bank,
			model.Payment.DeliveryCost,
			model.Payment.GoodsTotal,
			model.Payment.CustomFee,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, tx.Rebind(query), params...)
	if err != nil {
		return err
	}

	query, params, err = sq.Insert("models").
		Columns(
			"order_uid",
			"track_number",
			"entry",
			"delivery_id",
			"payment_id",
			"locale",
			"internal_signature",
			"customer_id",
			"delivery_service",
			"shard_key",
			"sm_id",
			"date_created",
			"oof_shard",
		).
		Values(
			model.OrderUID,
			model.TrackNumber,
			model.Entry,
			deliveryId,
			model.Payment.Transaction,
			model.Locale,
			model.InternalSignature,
			model.CustomerID,
			model.DeliveryService,
			model.Shardkey,
			model.SmID,
			model.DateCreated,
			model.OofShard,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, tx.Rebind(query), params...)
	if err != nil {
		return err
	}

	for _, item := range model.Items {
		query, params, err = sq.Insert("items").
			Columns(
				"chrt_id",
				"track_number",
				"price",
				"rid",
				"name",
				"sale",
				"size",
				"total_price",
				"nm_id",
				"brand",
				"status",
				"models_id",
			).
			Values(
				item.ChrtID,
				item.TrackNumber,
				item.Price,
				item.Rid,
				item.Name,
				item.Sale,
				item.Size,
				item.TotalPrice,
				item.NmID,
				item.Brand,
				item.Status,
				model.OrderUID,
			).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, tx.Rebind(query), params...)
		if err != nil {
			return err
		}

	}

	return tx.Commit()
}
