/*
	Ler uma temperatura em graus Fahrenheit, calcular e escrever o valor
	equivalente em graus Celsius.
 */

package main

import "fmt"

func main(){
	cel := 10.0
	fah := (cel * 9 / 5) + 32

	fmt.Printf("%.2f equivale à %.2f", cel, fah)
}
