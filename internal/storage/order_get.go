package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetOrders(ctx context.Context) ([]Order, error) {
	query, params, err := sq.Select(
		"m.order_uid",
		"m.track_number",
		"m.entry",
		"d.id",
		"d.name",
		"d.phone",
		"d.zip",
		"d.city",
		"d.address",
		"d.region",
		"d.email",
		"p.transaction",
		"p.request_id",
		"p.currency",
		"p.provider",
		"p.amount",
		"p.payment_dt",
		"p.bank",
		"p.delivery_cost",
		"p.goods_total",
		"p.custom_fee",
		"m.locale",
		"m.internal_signature",
		"m.customer_id",
		"m.delivery_service",
		"m.shard_key",
		"m.sm_id",
		"m.date_created",
		"m.oof_shard").
		From("models m").
		InnerJoin("deliveries d on d.id = m.delivery_id").
		InnerJoin("payments p on p.transaction=m.payment_id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var dest []Order
	err = s.db.SelectContext(ctx, &dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

//SELECT m.orderUID, m.trackNumber, m.entry, d.id, d.name, d.phone, d.zip, d.city, d.address, d.region, d.email, p.transaction, p.requestID, p.currency, p.provider, p.amount, p.paymentDt, p.bank, p.deliveryCost, p.goodsTotal, p.customFee, m.locale, m.internalSignature, m.customerID, m.deliveryService, m.shardkey, m.smID, m.dateCreated, m.oofShard FROM models m inner join deliveries d on d.id = m.delivery_id inner join payments p on p.transaction=m.payment_id;
