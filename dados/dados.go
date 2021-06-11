package dados

import "fmt"

func maps() {
	maquiagens := map[string]string{
		"BTM": "batom",
		"BSE": "base",
		"SMB": "sombra",
	}
	fmt.Println(maquiagens)
	base := maquiagens["BSE"]
	fmt.Println(base)

	maquiagens["BTM"] = "deliniador"
	fmt.Println(maquiagens)

	delete(maquiagens, "BSE")
	fmt.Println(maquiagens)
}

func slices() {
	maquiagens := []string{"batom", "base", "sombra"}
	fmt.Println(maquiagens)
	fmt.Println("len", len(maquiagens), "cap", cap(maquiagens))

	maquiagens = append(maquiagens, "rimel")
	fmt.Println(maquiagens)

	fmt.Println("len", len(maquiagens), "cap", cap(maquiagens))
}

func arrayComRange() {
	var maquiagens [5]string
	maquiagens[0] = "batom"
	maquiagens[1] = "base"
	maquiagens[2] = "sombra"

	fmt.Println(maquiagens)
	fmt.Println(len(maquiagens))

	for posicao, elemento := range maquiagens {
		fmt.Println("Maquiagem na posição", posicao, ":", elemento)
	}
}

func arrayComForTradicional() {
	var maquiagens [5]string
	maquiagens[0] = "batom"
	maquiagens[1] = "base"
	maquiagens[2] = "sombra"

	fmt.Println(maquiagens)
	fmt.Println(len(maquiagens))

	for posicao := 0; posicao < len(maquiagens); posicao++ {
		elemento := maquiagens[posicao]
		fmt.Println("Maquiagem na posição", posicao, ":", elemento)
	}
}
