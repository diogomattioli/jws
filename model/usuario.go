// Package modelsold contains the types for schema 'public'.
package model

// Usuario represents a row from 'public.usuario'.
type Usuario struct {
	ID    int    `json:"id"`    // id
	Nome  string `json:"nome"`  // nome
	Login string `json:"login"` // login
	Senha string `json:"senha"` // senha
	Ativo bool   `json:"ativo"` // ativo
	Grupo int    `json:"grupo"` // grupo
}
