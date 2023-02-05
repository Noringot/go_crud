package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func prepareResponseData(rowLen int) (row []sqlDataType, rowPtr []any) {
	row = make([]sqlDataType, rowLen)
	rowPtr = make([]any, rowLen)
	for i, _ := range row {
		rowPtr[i] = &row[i]
	}

	return row, rowPtr
}

func fmtResData(rowColumn []string, res [][]sqlDataType) ResponseDB {
	formatData := make([]map[string]any, len(res))
	for i, v := range res {
		mediate := map[string]any{}
		for k, val := range v {
			mediate[rowColumn[k]] = string(val)
		}
		formatData[i] = mediate
	}
	return formatData
}

func formatWhereParams(params []string) string {
	res := ""
	for i := 0; i < len(params); i += 2 {
		switch i {
		case len(params) - 2:
			res = res + fmt.Sprintf("%s='%s'", params[i], params[i+1])
		default:
			res = res + fmt.Sprintf("%s='%s' AND ", params[i], params[i+1])
		}
	}
	return res
}

func QueryContext(m *Model, queryString string) ResponseDB {
	rows, err := m.db.QueryContext(context.Background(), queryString)

	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	data := make([][]sqlDataType, 0)

	for rows.Next() {
		row, rowPtr := prepareResponseData(len(m.columns))
		if err := rows.Scan(rowPtr...); err != nil {
			log.Fatal(err)
		}
		data = append(data, row)
	}

	return fmtResData(m.columns, data)
}

func ExecContext(m *Model, queryString string) error {
	result, err := m.db.ExecContext(context.Background(), queryString)

	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		return errors.New(fmt.Sprintf("expected to affect 1 row, affected %d", rows))
	}
	return nil
}
