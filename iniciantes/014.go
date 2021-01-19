/*
	Um motorista de taxi deseja calcular o rendimento de seu carro na praça.
	Sabendo-se que o preço do combustível é de R$2,98 o litro, ler a marcação do
	odômetro no início e no fim do dia, o número de litros de combustível gasto e o
	valor recebido dos passageiros. Em seguida, calcular e escrever a média do
	consumo em Km/l e o lucro líquido do dia.
*/

package main

import "fmt"

func main(){
	valorCombustivel := 2.98
	odometroI := 0.0
	odometroF := 250.0
	litrosGastos := 21.0
	valorRecebido := 300.0

	gasto := litrosGastos * valorCombustivel
	mediaConsumo := (odometroF - odometroI) / litrosGastos
	lucroLiquido := valorRecebido - gasto

	fmt.Printf("Consumo médio: %.2fkm/l \nLucro Líquido: R$%.2f", mediaConsumo, lucroLiquido)
}