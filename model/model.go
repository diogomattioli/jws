package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

/*
	Id int `json:"id" gorm:"column:id_marca;primary_key"` // id
*/

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullFloat64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullBool struct {
	sql.NullBool
}

func (v NullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullBool) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Bool = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullTime struct {
	sql.NullTime
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}
	return nil
}

type ISearch interface {
	Search() []string
}

func Search(db *gorm.DB, obj interface{}, values url.Values) *gorm.DB {
	var fields []string
	if iface, ok := obj.(ISearch); ok {
		fields = iface.Search()
	} else {
		ty := reflect.TypeOf(obj).Elem()
		for i := 0; i < ty.NumField(); i++ {
			if ty.Field(i).Type.Name() == "string" || ty.Field(i).Type.Name() == "NullString" {
				fields = append(fields, ty.Field(i).Name)
			}
		}
	}

	if str := values.Get("search"); str != "" {
		for _, str := range strings.Split(str, " ") {
			for _, field := range fields {
				db = db.Or(fmt.Sprintf("%s ILIKE ?", toSnakeCase(field)), "%" + str + "%");
			}
		}
	}
	return db;
}

type IOrder interface {
	Order() []string
}

func Order(db *gorm.DB, obj interface{}, values url.Values) *gorm.DB {
	var fields []string
	if iface, ok := obj.(IOrder); ok {
		fields = iface.Order()
	} else {
		ty := reflect.TypeOf(obj).Elem()
		for i := 0; i < ty.NumField(); i++ {
			fields = append(fields, ty.Field(i).Name)
		}
	}

	if order := values.Get("order"); order != "" {
		for _, field := range fields {
			if strings.ToLower(field) == strings.ToLower(order) {
				return db.Order(fmt.Sprintf("%s ASC", toSnakeCase(order)))
			}
		}
	}
	return db;
}

func ObjModel(name string) (interface{}, interface{}) {
	switch name {
	case "arquivo":
		return new(Arquivo), new([]Arquivo)
	case "categoria":
		return new(Categoria), new([]Categoria)
	case "cep":
		return new(Cep), new([]Cep)
	case "cliente":
		return new(Cliente), new([]Cliente)
	case "comunicacao":
		return new(Comunicacao), new([]Comunicacao)
	case "endereco":
		return new(Endereco), new([]Endereco)
	case "equipamento":
		return new(Equipamento), new([]Equipamento)
	case "estoque":
		return new(Estoque), new([]Estoque)
	case "impressao":
		return new(Impressao), new([]Impressao)
	case "marca":
		return new(Marca), new([]Marca)
	case "ordem":
		return new(Ordem), new([]Ordem)
	case "ordemdesc":
		return new(Ordemdesc), new([]Ordemdesc)
	case "produto":
		return new(Produto), new([]Produto)
	case "usuario":
		return new(Usuario), new([]Usuario)
	default:
		return nil, nil
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}