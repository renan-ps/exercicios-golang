/*
	Ler um valor e escrever se é positivo, negativo ou zero.
*/

package main

import "fmt"

func main(){
	var numero = -3

	if numero > 0 {
		fmt.Println("Positivo")
	}else if numero < 0 {
		fmt.Println("Negativo")
	}else{
		fmt.Println("Zero")
	}
}