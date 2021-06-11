package lacosderepeticao

import "fmt"

func main() {
}

func forInfinito() {
	contador := 0
	for {
		fmt.Println("Contando", contador)
		contador++
	}
}

func forLikeWhile() {
	contador := 0
	for contador <= 10 {
		fmt.Println("Contando", contador)
		contador++
	}
}

func forTradicional() {
	for contador := 10; contador >= 0; contador-- {
		fmt.Println("Contando", contador)
	}
}
