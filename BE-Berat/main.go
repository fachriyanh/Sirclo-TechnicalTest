package main

import (
	"fmt"
	"net/http"

	"github.com/Sirclo-TechnicalTest/BE-Berat/controllers"
)

func main() {

	http.HandleFunc("/", controllers.Controller)

	fmt.Println("Server berjalan di Port 8080...")
	http.ListenAndServe(":8080", nil)
}
