package main

import "fmt"

// "src/math"

// var a string

type Carro struct {
	Nome string
}

func (c *Carro) andar() {
	c.Nome = "civic"
	fmt.Println(c.Nome, "andou")
}

func abc(a *int) {
	*a = 200
}

func oldmain() {

	carro := Carro{
		Nome: "a3",
	}

	carro.andar()

	fmt.Println(carro.Nome)

	// variavel := 10

	// abc(&variavel)

	// fmt.Println(variavel)

	// a := 10

	// fmt.Println(&a)

	// var ponteiro *int = &a

	// fmt.Println(*ponteiro)

	// carro := Carro{Nome: "corolla"}

	// carro.andar()

	// a = "victor"
	// fmt.Printf("%v \n%T", a, a)

	// resultado := math.Soma(1, 1)
	// fmt.Printf("\n%v", resultado)

	// res, err := http.Get("http://goodasdsadasdgle.com.br")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println(res.Header)

	// res, err := soma(11, 9)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println(res)

	// res, _ := soma(11, 9)
	// fmt.Println(res)

	// result := soma(11, 9)
	// fmt.Println(result)

	// result := somaTudo(11, 9, 1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	// resultado := func(x ...int) func() int {
	// 	result := 0

	// 	for _, v := range x {
	// 		result += v
	// 	}
	// 	return func() int {
	// 		return result * result
	// 	}
	// }

	// fmt.Println(resultado(11, 11, 11, 11)())
}

// func somaTudo(x ...int) int {
// 	result := 0

// 	for _, v := range x {
// 		result += v
// 	}

// 	return result
// }

// func soma(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func soma(x int, y int) (int, error) {
// 	res := x + y
// 	if res > 10 {
// 		return 0, errors.New("Total maior que 10")
// 	}

// 	return res, nil
// }
