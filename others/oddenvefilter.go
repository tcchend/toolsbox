/*
 数组中分离奇偶数
 */
package others

import "fmt"

//函数作为数据类型
type myFunc func(arr int) bool

//判断是奇数
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

//判断是偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

//根据函数来处理切片，实现奇偶数分组
func Filter(arr []int,f myFunc) []int {
	var result []int
	for _,value := range arr {
		if f(value) {
			result = append(result,value)
		}
	}
	return result
}

func main() {
	arr := []int{1,6,7,21,23,33,34,456,76,89,97,31,16,17,29}
	odd := Filter(arr,isOdd)
	fmt.Println("奇数有：",odd)
	even := Filter(arr,isEven)
	fmt.Println("偶数有：",even)
}