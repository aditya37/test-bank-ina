package mysql

import "context"

const mysqlQueryDeleteUser = `DELETE FROM mst_user WHERE id = ?`

func (m *mysqlClient) DeleteUserByID(ctx context.Context, userid int64) error {
	var (
		arg []interface{}
	)
	arg = append(arg, userid)
	if _, err := m.db.ExecContext(ctx, mysqlQueryDeleteUser, arg...); err != nil {
		return err
	}
	return nil
}
