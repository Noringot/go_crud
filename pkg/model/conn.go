package model

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

// Ping check database connection
func (m *Model) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	if err := m.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

// All Get all data from database, only register columns
// When you create instance of model, you need define columns what you want get from DB
// This clause accept for all methods below
func (m *Model) All() ResponseDB {
	columnsFormat := strings.Join(m.columns, ", ")
	queryString := fmt.Sprintf("SELECT %s FROM %s", columnsFormat, m.tb)

	return QueryContext(m, queryString)
}

// Where Get filtered data by multiple parameters
// Example: Model.Where("column", "value")
// Min two params. Required even number of params.
// Because odd arguments - title column in DB, even arguments - value this field
func (m *Model) Where(conditions ...string) ResponseDB {
	if len(conditions)%2 != 0 {
		log.Fatal("Must be quantity even params")
	}

	columnsFormat := strings.Join(m.columns, ", ")
	queryString := fmt.Sprintf("SELECT %s FROM %s WHERE %s", columnsFormat, m.tb, formatWhereParams(conditions))
	return QueryContext(m, queryString)
}

// Store Create a new record for current table
// When you create instance model, you fix this instance for concrete table in your DB
func (m *Model) Store(newRow map[string]string) error {
	columns := make([]string, 0)
	values := make([]string, 0)

	for _, v := range m.columns {
		if col, ok := newRow[v]; ok == true {
			columns = append(columns, v)
			values = append(values, fmt.Sprintf("'%s'", col))
		}
	}
	valuesFormat := strings.Join(values, ",")
	columnsFormat := strings.Join(columns, ", ")
	queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", m.tb, columnsFormat, valuesFormat)
	return ExecContext(m, queryString)
}

//func (m *Model) StoreWithReturn(newRow map[string]string) ResponseDB {}

// UpdateById Update record in table with id
func (m *Model) UpdateById(id string, newData map[string]string) error {
	values := make([]string, 0)

	for _, v := range m.columns {
		if opt, ok := newData[v]; ok == true {
			values = append(values, fmt.Sprintf("%s='%s'", v, opt))
		}
	}
	setOpt := strings.Join(values, ",")
	queryString := fmt.Sprintf("UPDATE %s SET %s WHERE id='%s'", m.tb, setOpt, id)
	return ExecContext(m, queryString)
}
func (m *Model) Delete(conditions ...string) error {
	if len(conditions)%2 != 0 {
		log.Fatal("Must be quantity even params")
	}

	queryString := fmt.Sprintf("DELETE FROM %s WHERE %s", m.tb, formatWhereParams(conditions))
	return ExecContext(m, queryString)
}
