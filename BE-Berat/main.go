package main

import (
	"BE-Berat/controllers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", controllers.Controller)

	fmt.Println("Server berjalan di Port 8080...")
	http.ListenAndServe(":8080", nil)
}
