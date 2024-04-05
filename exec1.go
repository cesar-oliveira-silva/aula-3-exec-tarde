package main

import "fmt"

// exercicio 1
type aluno struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

// type lista struct {
// 	student []aluno
// }

func main() {

	//exercicio 1
	var al aluno
	var paluno *aluno = &al
	//var alunos lista
	al.criaAluno("cesar", "silva", 36, "email@email.com", "senha")
	fmt.Println("Aluno criado:", al)

	mudarNomeP(&paluno.nome, &paluno.sobrenome, "augusto", "nonato")
	mudarIdade(&paluno.idade, 34)
	mudarEmail(&paluno.email, "emailnovo@email.com")
	mudarSenha(&paluno.senha, "654321")

	fmt.Println("Aluno alterado:", al)

}

// exercicio 1
func mudarNomeP(nome *string, sobrenome *string, novonome string, novosobrenome string) {
	*nome = novonome
	*sobrenome = novosobrenome
}

func mudarIdade(idade *int, novaidade int) {
	*idade = novaidade
}

func mudarEmail(email *string, novoemail string) {
	*email = novoemail
}

func mudarSenha(senha *string, novasenha string) {
	*senha = novasenha
}

func (a *aluno) criaAluno(name string, subname string, age int, mail string, pass string) {
	a.nome = name
	a.sobrenome = subname
	a.idade = age
	a.email = mail
	a.senha = pass
}
