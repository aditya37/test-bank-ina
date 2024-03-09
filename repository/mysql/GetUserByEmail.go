package mysql

import (
	"context"

	"github.com/adity37/task/model"
)

const mysqlQueryGetUserByEmail = `SELECT id,email FROM mst_user WHERE email = ?`

func (m *mysqlClient) GetUserByEmail(ctx context.Context, data model.RequestRegisterUser) (model.ResponseGetUserByID, error) {
	arg := []interface{}{}
	arg = append(arg, data.Email)

	var record model.ResponseGetUserByID
	if err := m.db.QueryRowContext(ctx, mysqlQueryGetUserByEmail, arg...).Scan(&record.UserID, &record.Email); err != nil {
		return model.ResponseGetUserByID{}, err
	}
	return record, nil
}
