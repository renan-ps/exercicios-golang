
/*
	Ler uma quantidade de chuva dada em polegadas, calcular e escrever o valor equivalente em milímetros.
	Sabe-se que 1" corresponde a 25,4mm.
 */

package main

import "fmt"

func main(){
	pol := 12.3
	mm  := pol * 25.4

	fmt.Printf("%.2f\" de chuva equivale à %.2fmm. ", pol, mm)
}
