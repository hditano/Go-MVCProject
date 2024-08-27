package main

import (
	"fmt"
	"hditano/MVCProject/initializers"
)

func main() {
	fmt.Println("Hello")
	initializers.ConnectToDatabase()
}
