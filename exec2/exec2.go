package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	email     string
	produtos  []produto
}

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

func main() {
	user1 := usuario{nome: "cesar", sobrenome: "silva", email: "email@email.com", produtos: []produto{}}
	fmt.Println("usuario criado: ", user1)
	var ponteiro *usuario = &user1
	fmt.Println("ponteiro usuario criado: ", ponteiro)

	prod1 := novoProduto("teste", 10.0)

	addProduto(ponteiro, prod1, 10)

	fmt.Println("produto adicionado via ponteiro: ", user1)

	deleteProd(ponteiro)

	fmt.Println("produto deletado via ponteiro: ", user1)
}

func novoProduto(name string, price float64) produto {
	return produto{nome: name, preco: price}
}

func addProduto(user *usuario, prod produto, qnt int) {
	prod.quantidade = qnt
	user.produtos = append(user.produtos, prod)
}

func deleteProd(user *usuario) {
	user.produtos = []produto{}
}
