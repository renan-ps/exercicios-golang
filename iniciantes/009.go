/*
	O custo ao consumidor de um carro novo é dado pelo custo de fábrica
	mais os percentuais do distribuidor e dos impostos (aplicados ao custo de
	fábrica). Supondo que o percentual do distribuidor seja de 28% e os impostos
	de 45%, ler o custo de fábrica de um carro e escrever o custo final ao
	consumidor.
 */
package main

import "fmt"

func main(){
	valor := 20000.0
	distribuidor := 28.0
	impostos := 45.0

	valorFinal := valor + (valor * (distribuidor / 100)) + (valor * (impostos / 100))
	fmt.Printf("Valor final: %.2f", valorFinal)
}
