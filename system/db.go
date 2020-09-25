package system

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	uri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", "localhost", "diogo", "jws3", "diogo")
	fmt.Println(uri)

	var err error
	DB, err = gorm.Open("postgres", uri)
	if err != nil {
		fmt.Println("failed to connect database", err)
		panic("failed to connect database")
	}
	//defer DB.Close()
	DB.LogMode(true)
	DB.SingularTable(true)
}
