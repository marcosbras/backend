package main

func soma(a int, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}
	return a + b, true
}

func main() {

	// Declaração de variáveis
	var a string
	a = "Hello World"

	// Declaração e inicialização de variáveis
	b := "Olá Mundo"

	// Declaração e inicialização de variáveis com tipo explícito e alteração de valor
	c := "Bonjour le monde"
	c = "Bonjour le monde, comment ça va?"

	resultado, ok := soma(10, 20)
	if ok {
		println("Resultado da soma:", resultado)
	} else {
		println("Erro: um dos números é negativo.")
	}

	println(a)
	println(b)
	println(c)
}
