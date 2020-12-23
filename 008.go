/*
	Ler o número de eleitores de um município e o número de votos brancos,
	nulos e válidos. Em seguida, calcular e escrever o percentual que cada tipo de
	voto representa em relação ao total de eleitores.
 */

package main

import "fmt"

func main(){
	eleitores := 100000.0
	brancos := 5000.0
	nulos := 7000.0
	validos := 88000.0

	pBrancos := (brancos / eleitores) * 100
	pNulos := (nulos / eleitores) * 100
	pValidos := (validos / eleitores) * 100

	fmt.Printf("Votos em branco: %.2f\n", pBrancos)
	fmt.Printf("Votos em branco: %.2f\n", pNulos)
	fmt.Printf("Votos em branco: %.2f\n", pValidos)


}