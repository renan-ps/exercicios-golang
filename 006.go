/*
	Ler um valor R correspondente ao raio de uma esfera, calcular e escrever
	o seu volume e a sua área.
 */

package main

import (
	"fmt"
	"math"
)

func main(){
	r := 15.0
	volume := (4 * math.Pi * math.Pow(r,3)) / 3
	area := 4 * math.Pi * math.Pow(r,2)

	fmt.Printf("Volume: %.2f\nÁrea: %.2f", volume, area)
}