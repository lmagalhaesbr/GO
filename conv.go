package main

import "fmt"

const ebulicaoK float64 = 373.15

func main() {
	var tempK float64 = ebulicaoK
	var tempC float64 = tempK - 273.15

	fmt.Printf("Ponto de ebulição da água:\n")
	fmt.Printf("%.2f graus Kelvin\n", tempK)
	fmt.Printf("%.2f graus Celsius\n", tempC)

}
