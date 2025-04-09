// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: description_attr.sql

package db

import (
	"context"
	"database/sql"
)

const createDescriptionAttr = `-- name: CreateDescriptionAttr :exec
INSERT INTO description_attr (
  description_attr_id, name, value, products_spu_id
) VALUES (
  ?, ?, ?, ?
)
`

type CreateDescriptionAttrParams struct {
	DescriptionAttrID string `json:"description_attr_id"`
	Name              string `json:"name"`
	Value             string `json:"value"`
	ProductsSpuID     string `json:"products_spu_id"`
}

func (q *Queries) CreateDescriptionAttr(ctx context.Context, arg CreateDescriptionAttrParams) error {
	_, err := q.db.ExecContext(ctx, createDescriptionAttr,
		arg.DescriptionAttrID,
		arg.Name,
		arg.Value,
		arg.ProductsSpuID,
	)
	return err
}

const deleteDescriptionAttr = `-- name: DeleteDescriptionAttr :exec
DELETE FROM description_attr
WHERE description_attr_id = ?
`

func (q *Queries) DeleteDescriptionAttr(ctx context.Context, descriptionAttrID string) error {
	_, err := q.db.ExecContext(ctx, deleteDescriptionAttr, descriptionAttrID)
	return err
}

const getDescriptionAttr = `-- name: GetDescriptionAttr :one
SELECT description_attr_id, name, value, products_spu_id FROM description_attr
WHERE description_attr_id = ? LIMIT 1
`

func (q *Queries) GetDescriptionAttr(ctx context.Context, descriptionAttrID string) (DescriptionAttr, error) {
	row := q.db.QueryRowContext(ctx, getDescriptionAttr, descriptionAttrID)
	var i DescriptionAttr
	err := row.Scan(
		&i.DescriptionAttrID,
		&i.Name,
		&i.Value,
		&i.ProductsSpuID,
	)
	return i, err
}

const listDescriptionAttrs = `-- name: ListDescriptionAttrs :many
SELECT description_attr_id, name, value, products_spu_id FROM description_attr
WHERE products_spu_id = ?
`

func (q *Queries) ListDescriptionAttrs(ctx context.Context, productsSpuID string) ([]DescriptionAttr, error) {
	rows, err := q.db.QueryContext(ctx, listDescriptionAttrs, productsSpuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DescriptionAttr
	for rows.Next() {
		var i DescriptionAttr
		if err := rows.Scan(
			&i.DescriptionAttrID,
			&i.Name,
			&i.Value,
			&i.ProductsSpuID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDescriptionAttr = `-- name: UpdateDescriptionAttr :exec
UPDATE description_attr
SET name = COALESCE(?, name),
    value = COALESCE(?, value),
    products_spu_id = COALESCE(?, products_spu_id)
WHERE description_attr_id = ?
`

type UpdateDescriptionAttrParams struct {
	Name              sql.NullString `json:"name"`
	Value             sql.NullString `json:"value"`
	ProductsSpuID     sql.NullString `json:"products_spu_id"`
	DescriptionAttrID string         `json:"description_attr_id"`
}

func (q *Queries) UpdateDescriptionAttr(ctx context.Context, arg UpdateDescriptionAttrParams) error {
	_, err := q.db.ExecContext(ctx, updateDescriptionAttr,
		arg.Name,
		arg.Value,
		arg.ProductsSpuID,
		arg.DescriptionAttrID,
	)
	return err
}
