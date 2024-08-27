package main

import (
	"fmt"
	"hditano/MVCProject/initializers"
)

func init() {
	initializers.ConnectToDatabase()
	initializers.SyncDB()
	initializers.AddData()
}

func main() {
	fmt.Println("Hello")
}
