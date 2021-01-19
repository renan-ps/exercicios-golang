/*
	Ler o salário mensal de uma pessoa e o percentual de reajuste, calcular e
	escrever o valor do salário reajustado.
*/

package main

import "fmt"

func main(){
	salario := 2000.0
	reajuste := 14.0

	novoSalario := salario + (salario * (reajuste / 100))

	fmt.Printf("Novo salário: R$%.2f", novoSalario)
}
