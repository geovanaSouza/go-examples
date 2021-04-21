package main

import "fmt"

func SemParametro() string {
	return "Exemplo de função sem parâmetro!"
}

func UmParametro(texto string) string {
	return texto
}

func DoisParametros(texto string, numero int) (string, int) {
	return texto, numero
}

// Variadic Function: Recebe número indeterminado de Parâmetros
func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func main() {
	fmt.Println(SemParametro())
	fmt.Println(UmParametro("Exemplo de função com um parâmetro!"))
	fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))

	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1))
	fmt.Println(Somando(1, 1, 1))
	fmt.Println(Somando(1, 1, 2, 4))
}
