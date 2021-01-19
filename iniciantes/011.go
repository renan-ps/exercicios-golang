/*
	O custo do seguro contra granizo numa comunidade típica de fazendeiros
	é 3,5% do valor de cobertura solicitado por acre, multiplicado pelo número de
	acres plantados. Supondo que as possibilidades de colheitas sejam limitadas a
	trigo, aveia e cevada, ler a cobertura desejada e o número de acres plantados
	para cada uma das três plantações e calcular e escrever o custo total do
	seguro.
 */

package main

import "fmt"

func main(){
	coberturaDesejada := 500.0
	trigo := 50.0
	aveia := 30.0
	cevada := 22.0

	total := (coberturaDesejada * 0.035) * (trigo + aveia + cevada)

	fmt.Printf("Valor do Seguro: R$%.2f", total)
}
