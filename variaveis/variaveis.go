package variaveis

import "fmt"

var (
	Variavel1  string = "Minha variavel"
	variavel2         = "Minha variavel"
	variavel3  string
	variavel9  float32
	variavel10 float64
	variavel11 bool
	variavel12 byte
)
var variavel4 string = "variavel 4"
var variavel5 = "variavel5"
var variavel6, variavel7, variavel8 string

func VariaveisLocais() {
	var variavel1 string = "variavel 1"
	var variavel2 = "variavel 2"
	variavel3 := "variavel 3"
	fmt.Println(variavel1, variavel2, variavel3)
	variavel1 = "novo valor variavel 1"
	variavel2 = "novo valor variavel 2"
	variavel3 = "novo valor variavel 3"
	fmt.Println(variavel1, variavel2, variavel3)
}

func constantesLocais() {
	const constante1 int = 1
	const constante2 = 2
	fmt.Println(constante1, constante2)
}

func grupoDeVariaveis() {
	var (
		variavel1 string = "variavel 1"
		variavel2        = "variavel 2"
	)
	fmt.Println(variavel1, variavel2)
}

func grupoDeConstantes() {
	const (
		constante1 string = "constante 1"
		constante2        = "constante 2"
	)
	fmt.Println(constante1, constante2)
}

func valoresPadrao() {
	println(variavel3)  // default value é vazio
	println(variavel9)  // default value é +0.000000e+000
	println(variavel10) // default value é +0.000000e+000
	println(variavel11) // default value é false
	println(variavel12) // default value é 0
}
