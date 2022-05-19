package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaInteiros(m map[string]int64) int64 {
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable is just to equality comparison
func Soma[T comparable](n1 T, n2 T) T {
	if n1 == n2 {
		return n1
	}
	return n2
}

func Max[T constraints.Ordered](n1 T, n2 T) T {
	if n1 > n2 {
		return n1
	}
	return n2
}

type stringer interface {
	String() string
}

type Mystring string

func (m Mystring) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}

	return result
}

func main() {
	// fmt.Println(SomaInteiros(map[string]int64{"a": 1, "b": 2, "c": 3}))
	// fmt.Println(SomaFloat(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.2}))

	fmt.Println(concat([]Mystring{"vic", "tor"}))

	// var x, y, z MyNumber
	// x = 1
	// y = 2
	// z = 3

	// fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))

	// fmt.Println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	// fmt.Println(SomaGenerica(map[string]float64{"a": 1.1, "b": 22.2, "c": 3.2}))
}
