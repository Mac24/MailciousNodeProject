package main

import (
	"MailciousNodeProject/utils"
	"fmt"
)

const (
	//转发指数
	FORWARD_INDEX = 0.80
	//处理延迟
	PROCESS_DELAY = 0.75
	//数据发送率
	DATA_TRABSMISSION_RATE = 0.75
	//邻居节点数
	NUMBER_NEIGHBOR_NODES = 0.8

	//相对速度
	RELATIVE_VELOCITY = 0.75
	//剩余能量
	RESIDUAL_ENERGY = 0.7
	//信号强度
	SIGNAL_STRENGTH = 0.75

	//数据可用性
	DATA_AVAILABILITY = 0.8
	//数据准确性
	DATA_ACCURACY = 0.7

	//判断阈值
	THERSHOLD = 0.5
	//判断个数
	NUM = 2

)

//信任要素初始判断矩阵(大类)
 var U0 = [3][3]int{{1,2,2}, {0,1,2}, {0,0,1}}
//权重(大类)
var S0 = []int64{5,3,1}

//第一小类，网络通信特性t1,U1,S1
var t1 = []float64{FORWARD_INDEX, PROCESS_DELAY, DATA_TRABSMISSION_RATE, NUMBER_NEIGHBOR_NODES}
var U1 = [4][4]int{{1,2,2,2}, {0,1,2,2}, {0,0,1,2}, {0,0,0,1}}
var S1 = []int64{7,5,3,1}

//第二小类,节点物理属性t2,U2,S2
var t2 = []float64{RELATIVE_VELOCITY, RESIDUAL_ENERGY, SIGNAL_STRENGTH}
var U2 = [3][3]int{{1,2,2}, {0,1,2}, {0,0,1}}
var S2 = []int64{5,3,1}

//第三小类,应用数据相关t3,U3,S3
var t3 = []float64{DATA_AVAILABILITY, DATA_ACCURACY}
var U3 = [2][2]int{{1,2},{0,1}}
var S3 = []int64{3,1}

func main() {
	var count int
	//计算大类的判断矩阵u,m0
	u0, m0 := utils.CalculationMatrix(S0)
	//计算大类的权重w0
	w0 := utils.CalculationWeight(m0)
	fmt.Printf("大类的判断矩阵u1:\n	%v \n大类的权重w1:\n	%v \n", u0, w0)


	//计算小类的判断矩阵u,m,权重w
	//第一小类
	u1, m1 := utils.CalculationMatrix(S1)
	w1 := utils.CalculationWeight(m1)
	//计算隶属度
	ms1 := utils.CalculationMembership(t1)
	fmt.Printf("第一小类的判断矩阵u1:\n	%v \n第一小类的权重w1:\n	%v \n第一小类的隶属度ms1:\n	%v \n", u1, w1, ms1)


	//第二小类
	u2, m2 := utils.CalculationMatrix(S2)
	w2 := utils.CalculationWeight(m2)
	ms2 := utils.CalculationMembership(t2)
	fmt.Printf("第二小类的判断矩阵u1:\n	%v \n第二小类的权重w1:\n	%v \n第二小类的隶属度ms1:\n	%v \n", u2, w2, ms2)


	//第三小类
	u3, m3 := utils.CalculationMatrix(S3)
	w3 := utils.CalculationWeight(m3)
	ms3 := utils.CalculationMembership(t3)
	fmt.Printf("第三小类的判断矩阵u1:\n	%v \n第三小类的权重w1:\n	%v \n第三小类的隶属度ms1:\n	%v \n", u3, w3, ms3)


	//计算一级模糊评判
	G := utils.GetArray(w1, w2, w3, ms1, ms2, ms3)
	fmt.Printf("一级模糊评判G:\n	%v \n", G)

	//计算二级模糊评判
	sum := utils.CalculationSecondJudge(w0, G)
	fmt.Printf("二级模糊评判sum:\n	%v \n", sum)

	for i:=0; i<len(sum); i++ {
		if sum[i] < THERSHOLD {
			count ++
		}
	}
	if count >= NUM {
		fmt.Printf("二级模糊评判数组中有%d个低于0.5的数值, 所以该节点是恶意节点!\n", count)
	} else {
		fmt.Printf("二级模糊评判数组中有%d个大于0.5的数值，所以该节点是正常节点!\n", count)
	}
}

