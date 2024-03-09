package mysql

import (
	"context"
	"errors"

	"github.com/adity37/task/model"
	sqldriver "github.com/go-sql-driver/mysql"
)

const mysqlQueryInsertUser = `INSERT INTO mst_user(name,email,password) VALUES(?,?,?)`

var ErrorDuplicate = errors.New("duplicate name or email")

func (m *mysqlClient) RegisterUser(ctx context.Context, data model.RequestRegisterUser) (int64, error) {
	var arg []interface{}
	arg = append(arg, data.Name, data.Email, data.Password)
	row, err := m.db.ExecContext(ctx, mysqlQueryInsertUser, arg...)
	if err != nil {
		n, ok := err.(*sqldriver.MySQLError)
		if !ok {
			return 0, err
		}
		if n.Number == 1062 {
			return 0, ErrorDuplicate
		}
	}
	return row.LastInsertId()
}
