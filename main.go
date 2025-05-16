package main

import (
	"fmt"
	"time"
)

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Cidade string 
	Estado string
	Cep string
}

func (c Cliente) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", c.Nome, c.Idade)
}

func main() {
	var p1 Cliente = Cliente{Nome: "deyvid marques"}
	var p2 Cliente = Cliente{Nome: "dvd marveira"}
	fmt.Println(p1.Nome)
	var p3 *Cliente = &p1 // definiu p3 como ponteiro de p1 passando como referência apontando o endereço.

	fmt.Println(&p1.Nome)
	p3.Nome = "David" // mudou o valor original de p1
	fmt.Println(p1.Nome)
	fmt.Println(p2.Idade)
	fmt.Println(&p3.Nome)

	multiplicacao()
}

func multiplicacao() { // essa forma de declarar função permite que ela acesse variáveis globais
	var fixo = 4
	multiplica := func(x int) int {
		return x * fixo
	}

	resultado := multiplica(5)
	fmt.Println(resultado)

	soma(2, 5)
}

func soma(a, b int) { // essa forma de declarar função não permite que ela acesse variáveis globais
	fmt.Println(a + b)

	clientes()
}

func clientes() {
	cliente1 := Cliente{
		Nome: "Siri",
		Idade: 19,
		Endereco: Endereco {
			Cidade: "Recife",
			Estado: "Pernambuco",
			Cep: "01010100",
		},
	}

	cliente2 := Cliente{
		Nome: "Iris",
		Idade: 19,
	}

	cliente2.Email = "iris@email.com"

	fmt.Println("Cliente 1:", cliente1)
	fmt.Println("Cliente 2:", cliente2.Nome)

	people()
}

func people() {
	var pessoas = map[string]int{}
	pessoas["dvd"] = 23
	pessoas["isis"] = 19

	if idade, ok := pessoas["dvd"]; ok {
			fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoa não existe no map")
	}

	nota()
}

func nota() { 
	nota := 50
	if nota >= 90 {
		fmt.Println("Aprovado com distinção")
	} else if nota >= 70 {
			fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

	players()
}

func players() {
	players := map[string]int{
		"david": 15,
	}

	if value, ok := players["david"]; ok {
		fmt.Println(value, "pontos")
	} else {
		fmt.Println("Jogador não encontrado")
	}

	switchCase()
}

func switchCase() {
	fmt.Println("Quando é sexta-feira?")
	today := time.Now().Weekday()

	switch time.Friday {
	case today + 0:
		fmt.Println("é hoje")
	case today + 1:
		fmt.Println("é amanhã")
	case today + 2:
		fmt.Println("é em dois dias")
	default:
		fmt.Println("Ainda está longe...")
	}

	loopFor()
}

// func loopFor() {
// 	sum := 0
// 	for i := 0; i < 10; i++{
// 		fmt.Println(i)
// 		sum += i
// 	}

// 	fmt.Println(sum)
// }

func loopFor() {
	sum := 0
	for sum < 20 {
		fmt.Println("loop")
		sum += 10
	}

	fmt.Println(sum)

	sliceNums()
}

func sliceNums() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
			fmt.Println(nums[i])
	}

	ranges()
}

func ranges() {
	nums := []string{"marveira", "marques", "oliveira"}
	for key, value := range nums{
		fmt.Println(key, value)
	}

	users := map[string]string{
		"nome": "Iris",
		"sexo": "feminino",
	}
	for key, value := range users {
		fmt.Println(key, value)
	}
}



