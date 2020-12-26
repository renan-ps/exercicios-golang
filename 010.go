/*
	Uma revendedora de carros usados paga a seus vendedores um salário
	fixo por mês, mais uma comissão, também fixa, para cada carro vendido e mais
	5% do valor das vendas efetuadas por eles. Ler o número de carros vendidos
	por um vendedor, o valor total de suas vendas, o salário fixo e o valor que ele
	recebe por carro. Em seguida, calcular e escrever o salário mensal do
	vendedor.
 */

package main

import "fmt"

func main(){
	carrosVendidos := 10.0
	totalVendas := 400000.0
	salarioFixo := 2500.0
	valorPorCarro := 100.0

	salarioFinal := salarioFixo + (valorPorCarro * carrosVendidos) + (totalVendas * 0.05)

	fmt.Printf("Salário total: R$%.2f", salarioFinal)
}