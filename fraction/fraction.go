package fraction

import (
	"fmt"
	"github.com/xxjwxc/public/mymath"
)

// FAL Fractional operation correlation
type FAL struct {
	Nume int64 // Nume numerator.分子
	Deno int64 // denominator (must not be zero).分母 (一定不为0)
}

// Format output
func (f FAL) String() string { // 格式化输出
	return fmt.Sprintf("%v/%v", f.Nume, f.Deno)
}

// Model Create a score (molecular, denominator) with a denominator default of 1
func Model(nd ...int64) FAL { // 创建一个分数(分子，分母)，分母默认为1
	var f FAL
	if len(nd) == 1 {
		f.Nume = nd[0]
		f.Deno = 1
	} else if len(nd) == 2 {
		f.Nume = nd[0]
		f.Deno = nd[1]
	}

	if f.Deno == 0 { // denominator is 0 .分母为0
		panic(fmt.Sprintf("fractional init error. denominator can't zero."))
	}

	return f
}

// Add Fraction addition
func (s *FAL) Add(f FAL) *FAL {
	// Getting the Minimum Common Multiplier 获取最小公倍数
	lcm := mymath.Lcm(f.Deno, s.Deno)
	s.broad(lcm)
	f.broad(lcm)

	s.Nume += f.Nume
	s.offset()
	return s
}

// Broadsheet  .阔张
func (s *FAL) broad(lcm int64) {
	s.Nume = s.Nume * (lcm / s.Deno)
	s.Deno = lcm
}

// Compression Finishing .压缩 整理
func (s *FAL) offset() {
	lcm := mymath.Gcd(s.Nume, s.Deno)

	s.Nume /= lcm
	s.Deno /= lcm
}

// Mul multiplication
func (s *FAL) Mul(f FAL) *FAL { // 乘法
	s.Deno *= f.Deno
	s.Nume *= f.Nume
	s.offset()
	return s
}

// Verdict Calculation results
func (s *FAL) Verdict() float64 { // 计算结果
	return float64(s.Nume) / float64(s.Deno)
}