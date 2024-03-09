package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_interface "github.com/adity37/task/repository/interface"
	"github.com/pressly/goose"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

type MYSQLConfig struct {
	Host              string
	Port              int
	Option            string
	Name              string
	User              string
	Password          string
	MaxConnection     int
	MaxIdleConnection int
}

type mysqlClient struct {
	db *sql.DB
}

func NewMYSQL(param MYSQLConfig) (_interface.DBReaderWriter, error) {
	connURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true&parseTime=true&loc=%s",
		param.User,
		param.Password,
		param.Host,
		param.Port,
		param.Name,
		"Asia%2FJakarta",
	)
	log.Printf(
		"MySQL Connection %s:%s@tcp(%s:%d)/%s",
		param.User,
		"********************",
		param.Host,
		param.Port,
		param.Name,
	)
	db, err := apmsql.Open("mysql", connURL)
	if err != nil {
		return nil, err
	}
	if param.MaxConnection > 0 {
		db.SetMaxOpenConns(param.MaxConnection)
	}
	if param.MaxIdleConnection > 0 {
		db.SetMaxIdleConns(param.MaxIdleConnection)
	}

	// migrator
	if err := goose.SetDialect("mysql"); err != nil {
		return nil, err
	}
	if err := goose.Up(db, "migration"); err != nil {
		return nil, err
	}

	return &mysqlClient{
		db: db,
	}, nil
}

func (m *mysqlClient) Close() error {
	return m.db.Close()
}
