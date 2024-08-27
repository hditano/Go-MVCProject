package main

import (
	"fmt"
	"hditano/MVCProject/initializers"
)

func init() {
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	fmt.Println("Hello")
}
