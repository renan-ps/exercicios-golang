/*
	Uma caixa de azulejos tem material suficiente para cobrir uma área de 1,5m².
	Ler as dimensões (comprimento, largura e altura) de uma cozinha
	retangular, calcular e escrever a quantidade de caixas de azulejos necessárias
	para cobrir todas as paredes. Considerar que não será descontada a área
	ocupada por portas e janelas.
*/

package main

import "fmt"

func main(){
	comprimento := 5.0
	largura := 3.0
	altura := 3.1

	area := comprimento * largura * altura
	caixas := area / 1.5

	fmt.Printf("Serão necessárias %.0f caixas de azulejos.", caixas)
}