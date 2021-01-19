package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/Sirclo-TechnicalTest/BE-Berat/controllers/routes"
)

func Controller(w http.ResponseWriter, r *http.Request) {

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	//declare html connection
	var getHtml, errGet = template.ParseFiles(wd + "/views/getData.html")
	if errGet != nil {
		log.Println(errGet.Error())
		return
	}

	var addHtml, errAdd = template.ParseFiles(wd + "/views/addData.html")
	if errAdd != nil {
		log.Println(errAdd.Error())
		return
	}

	var detailHtml, errDetail = template.ParseFiles(wd + "/views/detailData.html")
	if errDetail != nil {
		log.Println(errDetail.Error())
		return
	}

	switch r.Method {

	//Function to show data
	case "GET":

		aksi := r.URL.Query()["aksi"]

		switch len(aksi) {
		case 0:
			getHtml.Execute(w, routes.GetBerat())
		default:
			switch aksi[0] {
			case "tambah":
				addHtml.Execute(w, nil)
			case "detail":
				detailHtml.Execute(w, routes.DetailBerat(r.URL.Query()["tanggal"][0]))
			default:
				getHtml.Execute(w, routes.GetBerat())
			}
		}

	//Function to store data
	case "POST":

		var tanggal = r.FormValue("tanggal")
		max, err := strconv.Atoi(r.FormValue("max"))
		min, err := strconv.Atoi(r.FormValue("min"))
		var perbedaan = max - min

		if err != nil {
			log.Println(err)
		}

		var aksi = r.URL.Path
		switch aksi {
		case "/add":
			routes.AddBerat(tanggal, max, min, perbedaan)
		case "/detail":
			routes.DetailBerat(tanggal)
		default:
			getHtml.Execute(w, routes.GetBerat())
		}
		getHtml.Execute(w, routes.GetBerat())
	}

}
