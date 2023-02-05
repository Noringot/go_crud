package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Model struct {
	db      *sql.DB
	config  *mysql.Config
	tb      string
	columns []string
}

type ModelInterface interface {
	Init() *Model
	GetFillable()
}

func New(tableName string, columns []string) *Model {
	if len(columns) == 0 {
		log.Fatal("Expect columns name")
	}

	config := mysql.NewConfig()
	config.User = os.Getenv("DB_USERNAME")
	config.Passwd = os.Getenv("DB_PASSWORD")
	config.Addr = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_DATABASE")
	config.Net = "tcp"

	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return &Model{
		db:      db,
		tb:      tableName,
		config:  config,
		columns: columns,
	}
}

func (m *Model) Columns() []string {
	rows, err := m.db.QueryContext(context.Background(), fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = N'%s'", m.tb))
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	columns := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		columns = append(columns, name)
	}

	return columns
}
