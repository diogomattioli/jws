// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
	"time"
)

// Venda represents a row from 'public.venda'.
type Venda struct {
	ID        int              `json:"id"`        // id
	Cliente   int              `json:"cliente"`   // cliente
	Usuario   int              `json:"usuario"`   // usuario
	Nota      model.NullString `json:"nota"`      // nota
	Data      time.Time        `json:"data"`      // data
	Descricao model.NullString `json:"descricao"` // descricao
	Pagamento model.NullInt64  `json:"pagamento"` // pagamento
	Status    int16            `json:"status"`    // status
	Situacao  int16            `json:"situacao"`  // situacao
}

func (o *Venda) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}