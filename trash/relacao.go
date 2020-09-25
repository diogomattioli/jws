// Package modelsold contains the types for schema 'public'.
package trash

// Relacao represents a row from 'public.relacao'.
type Relacao struct {
	Estoque int `json:"estoque"` // estoque
	Produto int `json:"produto"` // produto
}