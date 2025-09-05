package main

import "fmt"

func sortTemp(a []float32) map[int][]float32 {
	res := make(map[int][]float32)
	for _, i := range a {
		k := int(i) / 10 * 10
		res[k] = append(res[k], i)
	}
	return res
}

func main() {
	//по примеру работает, но если использовать числа -10<x<10, то они будут приходить к 0
	a := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	b := sortTemp(a)
	fmt.Println(b)
}
