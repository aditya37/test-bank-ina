package mysql

import (
	"context"

	"github.com/adity37/task/model"
)

const mysqlQueryGetTaskByID = `SELECT 
		mt.id,
		mt.title,
		mt.description,
		mt.status,
		mt.modified_at,
		mt.created_at,
		mu.id AS user_id,
		mu.email
	FROM mst_task mt
	INNER JOIN mst_user mu ON mu.id = mt.user_id WHERE mt.id =?`

func (m *mysqlClient) GetTaskByID(ctx context.Context, data model.RequestGetTaskByID) (model.ResponseGetTaskById, error) {
	args := []interface{}{}
	args = append(args, data.TaskID)

	row := m.db.QueryRowContext(ctx, mysqlQueryGetTaskByID, args...)

	var record model.ResponseGetTaskById
	if err := row.Scan(
		&record.ID,
		&record.Title,
		&record.Description,
		&record.Status,
		&record.ModifiedAt,
		&record.CreatedAt,
		&record.UserTask.UserID,
		&record.UserTask.Email,
	); err != nil {
		return model.ResponseGetTaskById{}, err
	}
	return record, nil
}
