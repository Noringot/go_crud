package task

import (
	"github.com/Noringotq/go-crud/internal/pkg/dotenv"
	"github.com/Noringotq/go-crud/pkg/model"
	"log"
	"reflect"
)

type Fillable struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Text   string `json:"text"`
	IsDone int    `json:"is_done"`
}

var Task = Init()

func Init() *model.Model {
	dotenv.Load()
	m := model.New("go_crud", fmtColumns(Fillable{}))
	err := m.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func fmtColumns(t Fillable) []string {
	res := make([]string, 0)
	st := reflect.TypeOf(t)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if json, ok := field.Tag.Lookup("json"); ok {
			if json != "" {
				res = append(res, json)
			}
		}
	}
	return res
}
