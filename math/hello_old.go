package main

import (
	"fmt"
)

// declaração de variavel explicita.
var a string

func main() {
	/*a = "alex"
	//Declaração de variave implicita.
	nome := "Alex"
	fmt.Println(a)
	fmt.Println(nome)
	// Variaveis.
	/*a := 10
	b := "world"
	c := 3.144
	d := false
	e := `ouadlkfasdlf
	asdf
	asdf
	a
	sdf`
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)
	fmt.Printf("%T %v \n", e, e)
	resultado := math.SomaX(12)
	fmt.Printf("%T %v \n", resultado, resultado)
	fmt.Printf("%v \n", math.A)
	res, err := http.Get("http://adsfaeddgoogle.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v", res.Header)
	multi, err := math.Multiplica(5, 2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Total multiplicação: %v", multi)
	*/
	resultado := soma(1, 2)
	fmt.Println(resultado)
	resultado2, str := soma2(5, 2)
	fmt.Println(resultado2, str)
	resultado3 := soma3(5, 6)
	fmt.Println(resultado3)
	resultadoTudo := somatudo(2, 35, 8, 1, 89, 1)
	fmt.Println(resultadoTudo)

}

// função soma com retorno de um valor.
func soma(a int, b int) int {
	return a + b
}

// função soma com retorno de dois valroes.
func soma2(a int, b int) (int, string) {
	return a + b, "somou"
}

// função soma com retorno nomeado, não precisa colocar a variavel no return
func soma3(a int, b int) (result int) {
	result = a + b
	return
}

//Fuções variaticas.
// esta forma é como se fosse um array de inteiros na vairiavel x
func somatudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
