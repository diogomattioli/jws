package main

import (
	"jws/orm"
	"jws/rest"
)

func main() {
	orm.Init()
	defer orm.DB.Close()
	rest.Init()
}
