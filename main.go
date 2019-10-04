package main

import (
	"fmt"
)

const (
	Lx = 100
	Lt = 20

	hx  = 1
	hx2 = hx * hx
	ht  = 1
	ht2 = ht * ht

	T = int(Lt / ht)

	c  = 10
	c2 = c * c
)

func makeInitialsRo() []float64 {
	res := make([]float64, Lx)

	for i := 20; i <= 30; i++ {
		res[i] = 1
	}

	return res
}

func makeInitialsRoDt() []float64 {
	res := make([]float64, Lx)

	for i := 70; i <= 80; i++ {
		res[i] = 1
	}

	return res
}

func initialApprox(ro []float64, rodt []float64) []float64 {
	res := make([]float64, Lx)

	for i := 1; i < Lx-1; i++ {
		res[i] = (c2*(ro[i+1]-2*ro[i]+ro[i-1])/(2*hx2)+rodt[i]/ht)*ht2 + ro[i]
	}

	return res
}

func borderApprox(ro, ro1 []float64) []float64 {
	res := make([]float64, Lx)

	return res
}

func innerApprox(ro0, ro []float64) []float64 {
	res := make([]float64, Lx)

	for i := 1; i < Lx-1; i++ {
		res[i] = c2*(ro[i+1]-2*ro[i]+ro[i-1])/(2*hx2)*ht2 + 2*ro[i] - ro0[i]
	}

	return res
}

func main() {
	// начальные условия
	ro := makeInitialsRo()

	roDt := makeInitialsRoDt()

	fmt.Println(ro)
	// аппроксимация начальных условий
	ro1 := initialApprox(ro, roDt)

	// цикл аппроксимации
	for t := 0; t < T; t++ {
		ro := innerApprox(ro, ro1)
		ro, ro1 = ro1, ro
	}

	fmt.Println(ro1)
}
