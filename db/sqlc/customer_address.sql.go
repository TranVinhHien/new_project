// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: customer_address.sql

package db

import (
	"context"
	"database/sql"
)

const deleteCustomerAddress = `-- name: DeleteCustomerAddress :exec
DELETE FROM customer_address
WHERE id_address = ?
`

func (q *Queries) DeleteCustomerAddress(ctx context.Context, idAddress string) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerAddress, idAddress)
	return err
}

const getCustomerAddress = `-- name: GetCustomerAddress :one
SELECT id_address, customer_id, address, phone_number, create_date, update_date FROM customer_address
WHERE id_address = ? LIMIT 1
`

func (q *Queries) GetCustomerAddress(ctx context.Context, idAddress string) (CustomerAddress, error) {
	row := q.db.QueryRowContext(ctx, getCustomerAddress, idAddress)
	var i CustomerAddress
	err := row.Scan(
		&i.IDAddress,
		&i.CustomerID,
		&i.Address,
		&i.PhoneNumber,
		&i.CreateDate,
		&i.UpdateDate,
	)
	return i, err
}

const listCustomerAddresses = `-- name: ListCustomerAddresses :many
SELECT id_address, customer_id, address, phone_number, create_date, update_date FROM customer_address
WHERE customer_id = ?
`

func (q *Queries) ListCustomerAddresses(ctx context.Context, customerID string) ([]CustomerAddress, error) {
	rows, err := q.db.QueryContext(ctx, listCustomerAddresses, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CustomerAddress
	for rows.Next() {
		var i CustomerAddress
		if err := rows.Scan(
			&i.IDAddress,
			&i.CustomerID,
			&i.Address,
			&i.PhoneNumber,
			&i.CreateDate,
			&i.UpdateDate,
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

const listCustomerAddressesPaged = `-- name: ListCustomerAddressesPaged :many
SELECT id_address, customer_id, address, phone_number, create_date, update_date FROM customer_address
WHERE customer_id = ?
ORDER BY id_address
LIMIT ? OFFSET ?
`

type ListCustomerAddressesPagedParams struct {
	CustomerID string `json:"customer_id"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

func (q *Queries) ListCustomerAddressesPaged(ctx context.Context, arg ListCustomerAddressesPagedParams) ([]CustomerAddress, error) {
	rows, err := q.db.QueryContext(ctx, listCustomerAddressesPaged, arg.CustomerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CustomerAddress
	for rows.Next() {
		var i CustomerAddress
		if err := rows.Scan(
			&i.IDAddress,
			&i.CustomerID,
			&i.Address,
			&i.PhoneNumber,
			&i.CreateDate,
			&i.UpdateDate,
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

const updateCustomerAddress = `-- name: UpdateCustomerAddress :exec
UPDATE customer_address
SET address = COALESCE(?, address),
    phone_number = COALESCE(?, phone_number),
    update_date = NOW()
WHERE id_address = ?
`

type UpdateCustomerAddressParams struct {
	Address     sql.NullString `json:"address"`
	PhoneNumber sql.NullString `json:"phone_number"`
	IDAddress   string         `json:"id_address"`
}

func (q *Queries) UpdateCustomerAddress(ctx context.Context, arg UpdateCustomerAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateCustomerAddress, arg.Address, arg.PhoneNumber, arg.IDAddress)
	return err
}
