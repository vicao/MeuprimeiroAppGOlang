package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var nome string
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Olá, %s!\n", nome)
	var age int
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&age)
	fmt.Printf("Sua idade é: %d\n", age)
}
