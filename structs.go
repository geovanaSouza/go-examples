package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {

	p1 := Pessoa{nome: "Luísa"}
	fmt.Println(p1)

}
