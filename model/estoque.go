// Package modelsold contains the types for schema 'public'.
package model

// Estoque represents a row from 'public.estoque'.
type Estoque struct {
	ID         int        `json:"id"`         // id
	Titulo     string     `json:"titulo"`     // titulo
	Local      string     `json:"local"`      // local
	Codigo     NullString `json:"codigo"`     // codigo
	Preco      float32    `json:"preco"`      // preco
	Custo      float32    `json:"custo"`      // custo
	Quantidade int        `json:"quantidade"` // quantidade
	Pedido     int        `json:"pedido"`     // pedido
	Pedinte    NullString `json:"pedinte"`    // pedinte
	Descricao  NullString `json:"descricao"`  // descricao
	Ncm        NullString `json:"ncm"`        // ncm
	Ativo      bool       `json:"ativo"`      // ativo
}
