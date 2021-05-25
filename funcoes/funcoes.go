package funcoes

import "fmt"

func main() {
	// _, retorno2, retorno3 := FuncaoComDoisOuMaisRetornos()
	// fmt.Println(retorno2, retorno3)

	retorno1, retorno2, retorno3 := FuncaoComDoisOuMaisRetornosNomeados()
	fmt.Println(retorno1, retorno2, retorno3)
}

func funcaoPrivada() {
	fmt.Println("FuncaoSimplesPrivada")
}

func FuncaoPublica() {
	fmt.Println("FuncaoSimplesPublica")
}

func FuncaoComParametros(argumento1 string, argumento2 int) {
	fmt.Println("Primeiro argumento:", argumento1, "+ Segundo argumento:", argumento2)
}

func FuncaoComParametrosMesmoTipo(argumento1, argumento2 int, argumento3, argumento4 string) {
	fmt.Println("Primeiros argumentos:", argumento1, argumento2)
	fmt.Println("Segundos argumentos:", argumento3, argumento4)
}

func FuncaoComRetorno() int {
	return 1
}

func FuncaoComDoisOuMaisRetornos() (int, string, int) {
	retorno1 := 1
	retorno2 := "texto"
	retorno3 := 3
	return retorno1, retorno2, retorno3
}

func FuncaoComDoisOuMaisRetornosNomeados() (retorno1 int, retorno2 string, retorno3 int) {
	// seu codigo aqui fazendo alguma
	retorno1 = 1

	// seu codigo aqui fazendo outra coisa
	retorno2 = "texto"

	// seu outro codigo fazendo outra coisa
	retorno3 = 3

	return retorno1, retorno2, retorno3
}
