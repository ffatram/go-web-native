package main

import (
	"go-web-native-/config"
	"go-web-native-/controllers/categorycontroller"
	"go-web-native-/controllers/homecontroller"
	"go-web-native-/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()
	// Routes
	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("server running on port 3000")
	http.ListenAndServe(":3000", nil)

}
