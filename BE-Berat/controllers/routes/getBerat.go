package routes

import (
	"log"

	"github.com/Sirclo-TechnicalTest/BE-Berat/models"
)

func GetBerat() models.Response {

	var hasil []models.DataBerat
	var average []models.DataRata

	//query data total
	dataBerat, err := db.Query("select * from tableberat order by 1 desc")
	if err != nil {
		log.Println(err)
	}

	//read data total
	for dataBerat.Next() {
		var berat = models.DataBerat{}
		var err = dataBerat.Scan(&berat.Tanggal, &berat.Max, &berat.Min, &berat.Perbedaan)

		if err != nil {
			log.Println(err)
		}

		hasil = append(hasil, berat)
	}

	//query rata-rata
	dataRata, err := db.Query("select sum(max)/count(*), sum(min)/count(*), (sum(max)/count(*) - sum(min)/count(*) ) from tableberat order by 1 desc")
	if err != nil {
		log.Println(err)
	}

	//read rata-rata
	for dataRata.Next() {
		var avg = models.DataRata{}
		var err = dataRata.Scan(&avg.AvgMax, &avg.AvgMin, &avg.AvgPerbedaan)

		if err != nil {
			log.Println(err)
		}

		average = append(average, avg)
	}

	//response data
	return models.Response{
		Average: average,
		Data:    hasil,
	}

}
