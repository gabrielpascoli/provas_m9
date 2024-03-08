package main

import (
	"math/rand"
)

// simulando temperaturas de freazer e geladeira.
func Simulartemperatura() (float64, float64) {
	var1 := rand.Float64() * -15 + 10
	var2 := rand.Float64()*2 + 8
	return var1, var2
}
