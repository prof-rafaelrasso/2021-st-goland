package estrutura

import "fmt"

type MinhaEstrutura struct {
	atributoPrivado int
	AtributoPublico string
	naoInicializado string
}

func main() {
	estrutura := MinhaEstrutura{15, "texto", "batata"}
	estrutura2 := MinhaEstrutura{atributoPrivado: 15,
		AtributoPublico: "texto",
		naoInicializado: "batata"}

	estrutura3 := MinhaEstrutura{atributoPrivado: 15,
		AtributoPublico: "texto",
		naoInicializado: "batata3"}

	fmt.Println(estrutura == estrutura2)
	fmt.Println(estrutura == estrutura3)
	fmt.Println(&estrutura)
	fmt.Println(&estrutura2)
	fmt.Println(&estrutura == &estrutura2)
}
