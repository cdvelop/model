package model_test

import (
	"log"
	"testing"

	"github.com/cdvelop/model"
)

func TestGetTablesNames(t *testing.T) {

	obj := model.Object{
		Name:                "user",
		Table:               "user",
		PrincipalFieldsName: []string{},
		Fields: []model.Field{
			{Name: "name", Legend: "Nombre", NotRenderHtml: true},

			{Name: "cars", Legend: "Vehículos", SourceTable: "cars"},
			{Name: "other", Legend: "Vehículos", SourceTable: "other"},
		},
	}

	tables := obj.GetTablesNames()

	if len(tables) != 3 {
		log.Fatal("se esperaba solo los nombres de tablas y se obtuvo:", len(tables), tables)
	}
}
