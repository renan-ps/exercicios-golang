/*
	Ler os valores S1, S2 e S3 correspondentes aos comprimentos dos três
	lados de um triângulo, calcular e escrever a sua área.
 */

package main

import (
	"fmt"
	"math"
)

func main(){
	s1, s2, s3 := 3.0, 7.0, 9.0
	t := (s1 + s2 + s3) / 2
	area := math.Sqrt(t * (t-s1) * (t-s2) * (t-s3) )

	fmt.Printf("A área do triângulo é: %.2f", area)
}
