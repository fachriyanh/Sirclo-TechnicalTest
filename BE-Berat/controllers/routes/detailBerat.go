package routes

import (
	"log"

	"github.com/Sirclo-TechnicalTest/BE-Berat/models"
)

func DetailBerat(tanggal string) models.Response {

	var hasil []models.DataBerat

	//query data
	dataBerat, err := db.Query("select * from tableberat where tanggal=?", tanggal)
	if err != nil {
		log.Println(err)
	}

	defer dataBerat.Close()

	//read data
	for dataBerat.Next() {
		var berat = models.DataBerat{}
		var err = dataBerat.Scan(&berat.Tanggal, &berat.Max, &berat.Min, &berat.Perbedaan)

		if err != nil {
			log.Println(err)
		}

		hasil = append(hasil, berat)
	}

	//response data
	return models.Response{
		Data: hasil,
	}

}
