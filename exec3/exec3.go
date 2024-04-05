package main

import "fmt"

type produto struct {
	nome       string
	preco      float64
	quantidade float64
}

type servico struct {
	nome    string
	preco   float64
	minutos float64
}

type manutencao struct {
	nome  string
	preco float64
}

func main() {
	fmt.Println("exercicio 3")

	prod1 := produto{nome: "sabao", preco: 12.0, quantidade: 7}
	prod2 := produto{nome: "bucha", preco: 3.0, quantidade: 13}
	prod3 := produto{nome: "detergente", preco: 7.0, quantidade: 20}

	serv1 := servico{nome: "limpeza", preco: 120.0, minutos: 120}
	serv2 := servico{nome: "lavagem", preco: 30.0, minutos: 34}
	serv3 := servico{nome: "revisao", preco: 180.0, minutos: 90}

	manu1 := manutencao{nome: "conserto", preco: 150.0}
	manu2 := manutencao{nome: "prevencao", preco: 80.0}
	manu3 := manutencao{nome: "ajuste", preco: 70.0}

	var prods = []produto{}
	var servs = []servico{}
	var manuts = []manutencao{}

	prods = append(prods, prod1, prod2, prod3)
	servs = append(servs, serv1, serv2, serv3)
	manuts = append(manuts, manu1, manu2, manu3)

	totalProd := make(chan float64)
	totalServ := make(chan float64)
	totalManut := make(chan float64)

	go somaProdutos(prods, totalProd)
	go somaServicos(servs, totalServ)
	go somaManutencao(manuts, totalManut)

	fmt.Printf("A soma de produtos e igual a %.2f \n", <-totalProd)
	fmt.Printf("A soma de servicos e igual a %.2f\n", <-totalServ)
	fmt.Printf("A soma de manutencao e igual a %.2f\n", <-totalManut)
}
func somaProdutos(prods []produto, c chan float64) {
	fmt.Println("entrou soma produto")
	var soma float64
	for _, p := range prods {
		soma += p.preco * p.quantidade
	}
	fmt.Println("saiu soma produto")
	c <- soma

}

func somaServicos(serv []servico, c chan float64) {
	fmt.Println("entrou soma servicos")
	var soma float64
	for _, s := range serv {
		if s.minutos < 30 {
			soma += s.preco * 0.5
		} else {
			soma += s.preco * (s.minutos / 60)
		}
	}
	fmt.Println("saiu soma servicos")
	c <- soma

}

func somaManutencao(manut []manutencao, c chan float64) {
	fmt.Println("entrou soma manutencao")
	var soma float64
	for _, m := range manut {
		soma += m.preco
	}
	fmt.Println("saiu soma manutencao")
	c <- soma

}
