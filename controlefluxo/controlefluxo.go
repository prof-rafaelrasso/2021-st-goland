package controlefluxo

import "fmt"

// teste  de mesa de portas logicas
// true && true = true
// true || true = true
// false && true = false
// false || true = true
// true && false = false
// true || false = true
// false || false = false
// true && true && false = false
// true || true || false = true
// false || false || false = false
// ===========>
// (true || false) && false = false
//  ==> true && false = false

func main() {

}

func controleComSwitch() {
	fmt.Println("Escolha uma opção (0,1,2 ou 3)")

	var opcaoEscolhida int
	fmt.Scan(&opcaoEscolhida)

	switch opcaoEscolhida {
	case 0:
		fmt.Println("Voce eh falso")
	case 1, 2:
		fmt.Println("Voce eh verdadeiro")
	default:
		fmt.Println("Voce eh isentao")
	}
}

func ifElsio() {
	if ehVerdadeiro() {
		fmt.Println("Eh verdadeiro")
	} else if ehFalso() {
		fmt.Println("Eh falso")
	} else {
		fmt.Println("Elsizao")
	}
}

func ifTradicionalzao() {
	if ehVerdadeiro() {
		fmt.Println("Eh verdade emilio!")
	}

	if ehFalso() {
		fmt.Println("Eh falso")
	}
	
	if ehFalso() && ehVerdadeiro() {
		fmt.Println("Verdadeiro e Falso")
	}

	if ehFalso() || ehVerdadeiro() {
		fmt.Println("Verdadeiro ou Falso")
	}

	if ehVerdadeiro() || ehFalso() {
		fmt.Println("Verdadeiro ou Falso")
	}
}

func ehVerdadeiro() bool {
	fmt.Println("Verdadeiro!")
	return true
}

func ehFalso() bool {
	fmt.Println("Falso!")
	return false
}
