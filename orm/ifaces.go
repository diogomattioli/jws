package orm

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
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
				db = db.Or(fmt.Sprintf("%s ILIKE ?", toSnakeCase(field)), "%"+str+"%")
			}
		}
	}
	return db
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
	return db
}
