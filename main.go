package main

import (
	"fmt"
)

func kelvinToCelsius(kelvin float64) float64 {
	return (kelvin - 273);
}

func main() {

	var ebulicaoKelvin float64 = 373.2
	var ebulicaoCelsius = kelvinToCelsius(ebulicaoKelvin);


	fmt.Println("Temperatura de ebulição da água em Kelvin:", ebulicaoKelvin);
	fmt.Println("Temperatura de ebulição da água em Celsius:", ebulicaoCelsius);
}
