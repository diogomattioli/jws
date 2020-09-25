// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
)

// Importador represents a row from 'public.importador'.
type Importador struct {
	ID int `json:"id"` // id
	Nome string `json:"nome"` // nome
	Login string `json:"login"` // login
	Senha string `json:"senha"` // senha
	Marca []int `json:"marca"` // marca
	Ativo bool `json:"ativo"` // ativo
}

func (o *Importador) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}