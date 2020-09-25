package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	uri := fmt.Sprintf("host=%s password=%s user=%s dbname=%s", "localhost", "jws3", "diogo", "diogo")
	fmt.Println(uri)

	var err error
	DB, err = gorm.Open("postgres", uri)
	if err != nil {
		fmt.Println("failed to connect database", err)
		panic("failed to connect database")
	}
	DB.LogMode(true)
	DB.SingularTable(true)
}
