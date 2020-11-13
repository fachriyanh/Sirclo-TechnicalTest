package models

type DataBerat struct {
	Tanggal   string
	Max       int
	Min       int
	Perbedaan int
}

type DataRata struct {
	AvgMax       float64
	AvgMin       float64
	AvgPerbedaan float64
}

type Response struct {
	Average []DataRata
	Data    []DataBerat
}
