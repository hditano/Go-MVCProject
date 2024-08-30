package main

import (
	"hditano/MVCProject/initializers"
)

func init() {
	initializers.ConnectToDatabase()
	initializers.SyncDB()
	initializers.AddData()
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello World")
	//})

	// fmt.Print("\nStarting Server...\n")
	// http.ListenAndServe(":8080", nil)

	initializers.GetRecord()
}
