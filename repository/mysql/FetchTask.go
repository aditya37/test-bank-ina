package mysql

import (
	"context"

	"github.com/adity37/task/model"
)

const mysqlFetchTask = `SELECT user_id,id,description,status,title FROM mst_task`

func (m *mysqlClient) FetchTask(ctx context.Context) (model.ResponseFetchTask, error) {

	rows, err := m.db.QueryContext(ctx, mysqlFetchTask)
	if err != nil {
		return model.ResponseFetchTask{}, err
	}
	defer rows.Close()

	var data []model.TaskItem
	for rows.Next() {
		var record model.TaskItem
		if err := rows.Scan(
			&record.UserId,
			&record.TaskID,
			&record.Description,
			&record.Status,
			&record.Title,
		); err != nil {
			return model.ResponseFetchTask{}, err
		}
		data = append(data, record)
	}
	return model.ResponseFetchTask{
		Datas: data,
	}, nil
}
