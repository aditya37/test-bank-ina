package mysql

import (
	"context"

	"github.com/adity37/task/model"
)

const mysqlQueryGetUserByID = `SELECT id,name,email FROM mst_user WHERE id = ?`

func (m *mysqlClient) GetUserByID(ctx context.Context, data model.RequestGetUserByID) (model.ResponseGetUserByID, error) {
	var (
		record model.ResponseGetUserByID
		arg    []interface{}
	)
	arg = append(arg, data.UserID)

	err := m.db.
		QueryRowContext(ctx, mysqlQueryGetUserByID, arg...).
		Scan(
			&record.UserID,
			&record.Name,
			&record.Email,
		)
	if err != nil {
		return model.ResponseGetUserByID{}, err
	}
	return record, nil
}
