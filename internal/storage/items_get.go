package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetItems(ctx context.Context, modelId string) ([]Item, error) {
	query, params, err := sq.Select(
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
	).From("items").
		Where(sq.Eq{"models_id": modelId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var dest []Item

	err = s.db.SelectContext(ctx, &dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
