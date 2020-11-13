package main

import "ship/kapal"

func main() {
	var tictanic kapal.KapalPesiar

	tictanic.KodeKapal = "DJ2H"
	tictanic.Kapasitas = 1000
	tictanic.SpekMesin.KapasitasTanki = 1000

	kapal.InfoJarak(tictanic)
	//kapal.infoKode(tictanic) tidak bisa diakses

}
