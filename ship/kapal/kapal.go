package kapal

import "fmt"

type Kapal interface {
	GetJarakTempuh() int
	getDataKode() string
}

type PerahuMotor struct {
	KodeKapal string
	SpekMesin SpekMesin
}

type PerahuLayar struct {
	KodeKapal string
	Kapasitas int
	SpekMesin SpekMesin
}

type KapalPesiar struct {
	KodeKapal string
	Kapasitas int
	SpekMesin SpekMesin
}

type SpekMesin struct {
	BahanBakar     string
	KapasitasTanki int
}

func (perahumotor PerahuMotor) GetJarakTempuh() int {
	tanki := perahumotor.SpekMesin.KapasitasTanki
	return tanki / 2
}

func (perahulayar PerahuLayar) GetJarakTempuh() int {
	tanki := perahulayar.SpekMesin.KapasitasTanki
	kapasitas := perahulayar.Kapasitas
	var pembagi int
	switch perahulayar.SpekMesin.BahanBakar {
	case "bensin":
		pembagi = 200
	case "solar":
		pembagi = 500
	default:
		pembagi = 700
	}
	return tanki / (kapasitas / pembagi)
}

func (kapalpesiar KapalPesiar) GetJarakTempuh() int {
	tanki := kapalpesiar.SpekMesin.KapasitasTanki
	kapasitas := kapalpesiar.Kapasitas
	return tanki / (kapasitas / 1000)
}

func InfoJarak(kapal Kapal) {
	fmt.Println("Kapal anda bisa menempuh jarak", kapal.GetJarakTempuh())
}

func (perahumotor PerahuMotor) getDataKode() string {
	return perahumotor.KodeKapal
}

func (perahulayar PerahuLayar) getDataKode() string {
	return perahulayar.KodeKapal
}

func (kapalpesiar KapalPesiar) getDataKode() string {
	return kapalpesiar.KodeKapal
}

func infoKode(kapal Kapal) {
	fmt.Println("Kode Kapal Anda", kapal.getDataKode())
}
