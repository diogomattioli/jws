// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
)

// Notavendadesc represents a row from 'public.notavendadesc'.
type Notavendadesc struct {
	ID int `json:"id"` // id
	Titulo string `json:"titulo"` // titulo
	Preco float32 `json:"preco"` // preco
	Ncm string `json:"ncm"` // ncm
	Notavenda int `json:"notavenda"` // notavenda
	Venda int `json:"venda"` // venda
}

func (o *Notavendadesc) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}