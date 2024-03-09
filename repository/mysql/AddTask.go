package mysql

import (
	"context"

	"github.com/adity37/task/model"
)

const mysqlQueryInsertTask = `INSERT INTO mst_task(user_id,title,description,status) VALUES(?,?,?,?)`

func (m *mysqlClient) AddTask(ctx context.Context, data model.RequestAddTask) (model.ResponseAddTask, error) {
	args := []interface{}{}
	args = append(args, data.UserId, data.Title, data.Description, data.Status)

	row, err := m.db.ExecContext(ctx, mysqlQueryInsertTask, args...)
	if err != nil {
		return model.ResponseAddTask{}, err
	}
	id, _ := row.LastInsertId()
	return model.ResponseAddTask{
		ID: id,
	}, nil
}
