package mysql

import (
	"context"
	"fmt"

	"github.com/adity37/task/model"
)

const mysqlQueryUpdateUser = `UPDATE mst_user SET `

func (m *mysqlClient) UpdateUserByID(ctx context.Context, data model.RequestUpdateUser) error {
	var (
		arg   []interface{}
		query = mysqlQueryUpdateUser
		field string
	)

	if data.Email != "" {
		field = "email = ?"
		query = fmt.Sprintf("%s %s", query, field)
		arg = append(arg, data.Email)
	}

	if data.Name != "" {
		if field == "" {
			field = "name = ?"
			query = fmt.Sprintf("%s %s", query, field)
		} else {
			field = ",name = ?"
			query = fmt.Sprintf("%s %s", query, field)
		}
		arg = append(arg, data.Name)
	}

	if data.Password != "" {
		if field == "" {
			field = "password = ?"
			query = fmt.Sprintf("%s %s", query, field)
		} else {
			field = ",password = ?"
			query = fmt.Sprintf("%s %s", query, field)
		}
		arg = append(arg, data.Password)
	}

	// where
	query = fmt.Sprintf("%q WHERE id =?", query)
	arg = append(arg, data.UserID)

	if _, err := m.db.ExecContext(ctx, query, arg...); err != nil {
		return err
	}
	return nil
}
