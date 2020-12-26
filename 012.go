/*
	A potência necessária para iluminar adequadamente um cômodo é de 18 W/m2.
	Ler a potência de uma lâmpada e as dimensões (largura e comprimento)
	do cômodo, calcular e escrever o número de lâmpadas necessárias para a
	iluminação.
 */

package main

import "fmt"

func main(){
	potencia := 120.0
	largura := 10.0
	comprimento := 4.0

	lampadas := (18.0 / potencia) * (largura * comprimento)

	fmt.Printf("Lampadas necessárias: %.0f", lampadas)
}