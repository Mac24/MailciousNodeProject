package utils

import (
	"MailciousNodeProject/fraction"
	"math"
)

//计算判断矩阵U, mi
func CalculationMatrix(Si []int64) (interface{}, []fraction.FAL) {
	//判度矩阵U(不包括mi,wi)
	var u [][]interface{}
	//判断矩阵中的mi
	var mi []fraction.FAL
	for i:=0; i < len(Si); i++ {
		var ret fraction.FAL
		var tmp_mul *fraction.FAL
		var a []interface{}
		var tmp int64
		fValue := fraction.Model(1)
		for j:=0; j < len(Si); j++{
			ret_mul := fraction.Model(1)
			if Si[i] >= Si[j]{
				//公式(1)
				tmp = (Si[i] - Si[j] + 1)
				ret = fraction.Model(tmp)
			} else {
				//公式(2)
				//b =  Tools(float64(1/float64(Si[j]-Si[i]+1)))
				tmp = Si[j]-Si[i]+1
				ret = fraction.Model(1, tmp)
			}
			tmp_mul = fValue.Mul(ret)
			//fmt.Println("tmp_mul:", tmp_mul, "ret:", ret)
			ret_mul = *ret_mul.Mul(*tmp_mul)
			//fmt.Println("ret_mul:", ret_mul)
			if ret.Deno == 1 {
				a = append(a, ret.Verdict())
			} else {
				a = append(a, ret)
			}

		}
		mi = append(mi, *tmp_mul)
		//fmt.Println("mi:", mi)
		u = append(u, a)
	}
	return u, mi
}

//计算判断矩阵中的权重wi
func CalculationWeight(Mi []fraction.FAL) [][]float64 {
	//fmt.Println("Mi:", Mi)
	var ci, sum float64
	//判断矩阵中的wi
	var wi, tmp_ci []float64
	var w [][]float64

	for _, v := range Mi {
		//计算ci
		tmp_v := v.Verdict()
		ci = math.Pow(tmp_v, 1.0/float64(len(Mi)))
		sum += ci
		//fmt.Println("tmp_v:", tmp_v,"len(Mi):", len(Mi), "ci:",ci, "float64(1/len(Mi)):", 1.0/float64(len(Mi)), "sum:",sum)
		tmp_ci = append(tmp_ci, ci)
	}
	for _, v := range tmp_ci {
		var weight float64
		//计算wi
		weight = (v / sum)
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