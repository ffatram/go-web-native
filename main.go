package main

import (
	"go-web-native-/config"
	"go-web-native-/controllers/categorycontroller"
	"go-web-native-/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("server running on port 3000")
	http.ListenAndServe(":3000", nil)

}
