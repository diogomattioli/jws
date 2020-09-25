// Package modelsold contains the types for schema 'public'.
package model

// Arquivo represents a row from 'public.arquivo'.
type Arquivo struct {
	ID       int    `json:"id"`       // id
	Ordem    int    `json:"ordem"`    // ordem
	Arquivo  string `json:"arquivo"`  // arquivo
	Mimetype string `json:"mimetype"` // mimetype
}