package utils

import (
	"math"
)

//计算判断矩阵U, mi
func CalculationMatrix(Si []int) ([][]float64, []float64) {
	//判度矩阵U(不包括mi,wi)
	var u [][]float64
	var b float64
	//判断矩阵中的mi
	var mi []float64
	for i:=0; i < len(Si); i++ {
		var a []float64
		var tmp = 1.0
		for j:=0; j < len(Si); j++{
			if Si[i] >= Si[j]{
				//公式(1)
				b = Tools(float64(Si[i] - Si[j] + 1))
			} else {
				//公式(2)
				b =  Tools(float64(1/float64(Si[j]-Si[i]+1)))
			}
			tmp *= b
			a = append(a, b)
		}
		u = append(u, a)
		mi = append(mi, tmp)
	}
	return u, mi
}

//计算判断矩阵中的权重wi
func CalculationWeight(Mi []float64) [][]float64 {
	var ci, sum float64
	//判断矩阵中的wi
	var wi []float64
	var w [][]float64

	for _, v := range Mi {
		var weight float64
		//计算ci
		ci = math.Pow(v, 1/float64(len(Mi)))
		sum += ci
		//计算wi
		weight = Tools(ci / sum)
		wi = append(wi, weight)
	}
	w = append(w, wi)
	return w
}

//计算隶属度
func CalculationMembership(b []float64) [][]float64 {
	var lineArray []float64
	var tmpArray [][]float64
	for _, value := range b {
		lineArray = CalculationSum(value)
		tmpArray = append(tmpArray, lineArray)
	}
	return tmpArray
}

//计算一级模糊评判
func CalculationFistJudge(wi [][]float64, ri [][]float64) []float64 {
	var gi []float64
	//矩阵相乘计算
	for i:=0; i<len(wi); i++ {
		for j:=0; j<len(ri[i]); j++ {
			var tmp float64
			for k:=0; k<len(ri);k++ {
				tmp += wi[i][k] * ri[k][j]
				tmp = Tools(tmp)
				//fmt.Println(tmp)
			}
			gi = append(gi, tmp)
		}
	}
	return gi
}

func GetArray(w1,w2,w3 [][]float64, r1, r2, r3 [][]float64) [][]float64 {
	var Gi [][]float64
	g1 := CalculationFistJudge(w1, r1)
	g2 := CalculationFistJudge(w2, r2)
	g3 := CalculationFistJudge(w3, r3)

	Gi = append(Gi, g1, g2, g3)
	return Gi
}


//计算二级模糊评判
func CalculationSecondJudge(w0 [][]float64, Gi [][]float64) []float64 {
	var sumArray []float64
	for i:=0; i<len(w0); i++ {
		for j:=0; j<len(Gi[i]); j++ {
			var tmp float64
			for k:=0; k<len(Gi); k++ {
				tmp += Tools(w0[i][k]*Gi[k][j])
			}
			sumArray = append(sumArray, tmp)
		}
	}
	return sumArray
}