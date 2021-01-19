/*
	Ler um valor inteiro e escrever o seu antecessor e o seu sucessor.
 */

package main

import "fmt"

func main(){
	x := 10
	suc := x + 1
	ant := x - 1

	fmt.Printf("Antecessor: %v\nSucessor: %v", suc, ant)
}

