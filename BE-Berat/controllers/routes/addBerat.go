package routes

import (
	src "BE-Berat/config"
	"BE-Berat/models"
	"log"
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
