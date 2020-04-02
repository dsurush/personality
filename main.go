package main

import (
	"MF/db"
	"fmt"
)

func main() {
	db.InitDb()
	fmt.Println("Init Successfully")

}
