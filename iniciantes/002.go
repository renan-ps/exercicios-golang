/*
	Ler o nome de uma pessoa na forma “nome” seguido por “sobrenome” e escrever na forma “sobrenome” seguido por “nome”.
	Exemplo:
	entrada: “Fulano”, “de Tal”
	saída: “de Tal”, “Fulano”
 */

package main

import "fmt"

func main(){
	nome, sobrenome := "Renan", "Pereira"

	fmt.Println(sobrenome, nome)
}
