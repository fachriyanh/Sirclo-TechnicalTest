package routes

import (
	"log"

	"github.com/Sirclo-TechnicalTest/BE-Berat/models"

	src "github.com/Sirclo-TechnicalTest/BE-Berat/config"
)

//db connection
var db, err = src.Connect()

func AddBerat(tanggal string, max int, min int, perbedaan int) models.Response {

	//query data
	_, err = db.Exec("delete from tableberat where tanggal=?", tanggal)
	_, err = db.Exec("insert into tableberat values (?, ?, ?, ?)", tanggal, max, min, perbedaan)
	if err != nil {
		log.Println(err)
	}

	//response data
	return models.Response{
		Data: []models.DataBerat{},
	}
}
