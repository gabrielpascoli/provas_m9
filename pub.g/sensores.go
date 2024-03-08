package main

import (
	"math/rand"
)

// simularLeituraSensor simula uma leitura do sensor SPS30, retornando dois valores, PM2.5 (µg/m³) e PM10 (µg/m³).
func SimularSPS30() (float64, float64) {
	var1 := rand.Float64() * 10
	var2 := rand.Float64()*10 + 10
	return var1, var2
}
