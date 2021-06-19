package main

import "fmt"

type MeuTipo struct {
	arg1 int
	arg2 string
	arg3 string
}

func main() {
	meuSegundoTipo := MeuTipo{1, "argumento2", "argumento3"}

	fmt.Println(meuSegundoTipo)

	meuSegundoTipo.semPonteiro("qualquercoisa")

	fmt.Println(meuSegundoTipo)

	meuSegundoTipo.comPonteiro("qualquer coisa com ponteiro agora")

	fmt.Println(meuSegundoTipo)

}

func (m *MeuTipo) comPonteiro(algumaCoisa string) string {
	m.arg3 = algumaCoisa
	return algumaCoisa
}

func (m MeuTipo) semPonteiro(algumaCoisa string) string {
	m.arg3 = algumaCoisa
	return algumaCoisa
}
