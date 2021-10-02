package example

import "fmt"

func init() {
	fmt.Println("run init func")
}

func GetDB() {
	defer closeDB()
	fmt.Println("DB is online")
}

func closeDB() {
	fmt.Println("DB is offline")
}
