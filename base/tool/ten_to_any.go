package tool

import (
	"math"
	"strconv"
	"strings"
)

var tenToAnyMap map[int]string = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z"}

//10进制转任意进制，余数append
func TenToAny(num, n int) string {
	new_num_str :="" //转换后的字符串
	var remainder int //余数
	var remainder_string string  //余数字符串表示
	for num != 0{   //num不为0
		remainder=num%n    //计算余数
		if 76>remainder &&remainder>9{
			remainder_string=tenToAnyMap[remainder]
		}else{
			remainder_string=strconv.Itoa(remainder)//数字的字符串形式
		}
		new_num_str=remainder_string+new_num_str //拼接字符串
		num=num/n
	}
	return new_num_str
}

func findkeyMap(in string) int {
	result :=-1
	for k,v :=range tenToAnyMap{
		if in ==v{
			result=k
		}
	}
	return result
}

//转十进制，计算权重和
func AntYoTen(num string, n int) int {
	var newNum float64//转换后的十进制数
	newNum=0.0
	nNum :=len(strings.Split(num,""))-1
	for _,value:=range strings.Split(num,""){
		tmp :=float64(findkeyMap(value))
			if tmp !=-1{//找到键值
				newNum = newNum + tmp*math.Pow(float64(n), float64(nNum)) // 利用权重计算十进制数
				nNum = nNum - 1 // 更新权重
			}else{
				break
			}
		}
	return int(newNum)
}






























