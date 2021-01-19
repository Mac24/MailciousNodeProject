package utils

import (
	"fmt"
	"strconv"
)

//计算公式
func CalculationFormula1(v float64) float64 {
	var tmp float64
	if v <= 0.2 {
		tmp = 1.0
	} else if (v > 0.2 && v <= 0.4) {
		tmp = (v-0.4)/(0.2-0.4)
	} else {
		tmp = 0
	}
	return tmp
}

func CalculationFormula2(v float64) float64 {
	var tmp float64
	if v <= 0.2 {
		tmp = 0
	} else if (v > 0.2 && v <= 0.4) {
		tmp = (v-0.2)/(0.4-0.2)
	} else if (v > 0.4 && v <= 0.6){
		tmp = (v-0.6)/(0.4-0.6)
	} else {
		tmp = 0
	}
	return tmp
}

func CalculationFormula3(v float64) float64 {
	var tmp float64
	if v <= 0.4 {
		tmp = 0
	} else if (v > 0.4 && v <= 0.6) {
		tmp = (v-0.4)/(0.6-0.4)
	} else if (v > 0.6 && v <= 0.8){
		tmp = (v-0.8)/(0.6-0.8)
	} else {
		tmp = 0
	}
	return tmp
}

func CalculationFormula4(v float64) float64 {
	var tmp float64
	if v <= 0.6 {
		tmp = 0
	} else if (v > 0.6 && v <= 0.8) {
		tmp = (v-0.6)/(0.8-0.6)
	} else {
		tmp = 1.0
	}
	return tmp
}

func CalculationSum(v float64) []float64 {
	var tmp []float64
	a := Tools(CalculationFormula1(v))
	b := Tools(CalculationFormula2(v))
	c := Tools(CalculationFormula3(v))
	d := Tools(CalculationFormula4(v))
	tmp = append(tmp, a, b, c, d)
	return tmp
}

func Tools(v float64) float64 {
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
	return v
}