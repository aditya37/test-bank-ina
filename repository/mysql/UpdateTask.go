package mysql

import (
	"context"
	"fmt"

	"github.com/adity37/task/model"
)

const mysqlQueryUpdateTask = `UPDATE mst_task SET `

func (m *mysqlClient) UpdateTask(ctx context.Context, data model.RequestUpdateTask) error {
	var (
		arg   []interface{}
		query = mysqlQueryUpdateTask
		field string
	)

	// update title
	if data.Title != "" {
		field = "title = ?"
		query = fmt.Sprintf("%s %s", query, field)
		arg = append(arg, data.Title)
	}

	// desc
	if data.Description != "" {
		if field == "" {
			field = "description = ?"
			query = fmt.Sprintf("%s %s", query, field)
		} else {
			field = ",description = ?"
			query = fmt.Sprintf("%s %s", query, field)
		}
		arg = append(arg, data.Description)
	}

	// status
	if data.Status != "" {
		if field == "" {
			field = "status = ?"
			query = fmt.Sprintf("%s %s", query, field)
		} else {
			field = ",status = ?"
			query = fmt.Sprintf("%s %s", query, field)
		}
		arg = append(arg, data.Status)
	}

	// where clause
	arg = append(arg, data.TaskID)
	query = fmt.Sprintf("%s WHERE id = ?", query)

	if _, err := m.db.ExecContext(ctx, query, arg...); err != nil {
		return err
	}
	return nil
}
