package mysql

import (
	"context"

	"github.com/adity37/task/model"
)

const mysqlQueryDeleteTask = `DELETE FROM mst_task WHERE id = ?`

func (m *mysqlClient) DeleteTaskByID(ctx context.Context, data model.RequestDeleteTask) error {
	arg := []interface{}{}
	arg = append(arg, data.TaskID)

	if _, err := m.db.ExecContext(ctx, mysqlQueryDeleteTask, arg...); err != nil {
		return err
	}
	return nil
}
