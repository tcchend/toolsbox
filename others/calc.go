/*
  用栈实现复杂算术表达式
 */

package others

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int     //栈最大可以存放个数
	Top    int     //栈顶
	arr    [50]interface{} //模拟栈
}

func (this *Stack) Push(val interface{}) (err error) {
	if this.IsFull() {
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *Stack) Pop() (val interface{}, err error) {
	if this.IsEmpty() {
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func (this *Stack) List() {
	if this.IsEmpty()  {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func (this *Stack) IsEmpty() bool {
	if this.Top == -1 {
		return true
	}
	return false
}

func (this *Stack) IsFull() bool {
	if this.Top == this.MaxTop-1  {
		return true
	}
	return false
}

//判断一个字符是不是一个运算符【+，-，*，/,(,)】
func (this *Stack) IsOper(val int) bool {
	if val == 40 || val == 41 || val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) Cal(num1 float64, num2 float64, oper int) (res float64) {
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误",oper)
	}
	return res
}

//优先级定义
func (this *Stack) Priority(oper int) int {
	switch (oper) {
	//case 41:
	//	return 4
	case 42, 47:
		return 3
	case 43, 45:
		return 2
	case 40:
		return 1
	default:
		return 0
	}
}

func (this *Stack) isDigitNum(c byte) bool {
	return '0' <= c && c <= '9' || c == '.' || c == '_' || c == 'e' || c == '-' || c == '+'
}

func ExpressionConver(exp string,expLen int) float64{
	numStack := &Stack{ //数栈
		MaxTop: expLen,
		Top:    -1,
	}
	operStack := &Stack{ //符号栈
		MaxTop: expLen,
		Top:    -1,
	}
	//定义一个索引，便于扫描
	index := 0
	num1,num2 := 0.0,0.0
	result := 0.0
	oper := 0
	keepNum := ""
	for {
		//处理多位数思路1、定义一个变量keepNum string 做拼接 2、每次要向index的前面字符测试，看看是不是运算符，然后处理
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0]) //字符对应的ASCIL碼
		if operStack.IsOper(temp) {
			//如果operStack是一个空栈，直接入栈,或者'('直接入栈
			if operStack.IsEmpty() || temp == 40 {
				operStack.Push(temp)
			} else if temp == 41 { // )
				for {
					if operStack.IsEmpty() {
						fmt.Println("表达式有问题")
						break
					}
					tempOper, _ := operStack.Pop()
					oper = tempOper.(int)
					if oper != 40 {
						tempNum1, _ := numStack.Pop()
						tempNum2, _ := numStack.Pop()
						num1 = tempNum1.(float64)
						num2 = tempNum2.(float64)
						result = operStack.Cal(num1, num2, oper) //将计算结果重新入栈
						numStack.Push(result)
					} else {
						break
					}

				}

			} else if operStack.Priority(operStack.arr[operStack.Top].(int)) >= operStack.Priority(temp) {
				//如果发现operStack栈的运算符的优先级大于等于当前准备入栈的运算符的优先级， 就从符号栈pop出，并从数字栈也pop出两个数，进行运算
				//运算后的结果再重新入栈到数栈，符号再入符号栈
				tempNum1, _ := numStack.Pop()
				tempNum2, _ := numStack.Pop()
				tempOper, _ := operStack.Pop()
				num1 = tempNum1.(float64)
				num2 = tempNum2.(float64)
				oper = tempOper.(int)
				result = operStack.Cal(num1, num2, oper)
				//将计算结果重新入栈
				numStack.Push(result)
				//当前符号压入符号栈
				operStack.Push(temp)
			} else {
				operStack.Push(temp)
			}

		} else { //说明是数
			//处理多位数思路1、定义一个变量keepNum string 做拼接 2、每次要向index的后面字符测试，看看是不是运算符，然后处理,如果已经到表达式最后，直接将keepNum
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseFloat(keepNum,  64)
				numStack.Push(val)
			} else {
				//向index后面测试看看是不是运算符
				if operStack.IsOper(int([]byte(exp[index+1:index+2])[0])) {
					val, _ := strconv.ParseFloat(keepNum,  64)
					numStack.Push(val)
					keepNum = ""
				}

			}
			//val,_ := strconv.ParseInt(ch,10,64)
			//numStack.Push(int(val))
		}
		//继续扫描,先判断index是否已经扫描到表达式最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	//如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈去除两个数字,运算后的结果入数栈，直至符号栈为空
	for {
		if operStack.IsEmpty() {
			break // 退出条件
		}
		tempNum1, _ := numStack.Pop()
		tempNum2, _ := numStack.Pop()
		tempOper, _ := operStack.Pop()
		num1 = tempNum1.(float64)
		num2 = tempNum2.(float64)
		oper = tempOper.(int)
		result = operStack.Cal(num1, num2, oper)
		//将计算结果重新入栈
		numStack.Push(result)
	}
	//如果我们运算没有问题，表达式也正确，则结果就是unmStack最后数
	res, _ := numStack.Pop()
	return res.(float64)

}

func Demo()  {
	exp := "(100*0.95+10*0.05)*10/100*64/100+(10*0.4+20*0.4+30*0.2)*16/100+10*0.025+10*0.025"
	res := ExpressionConver(exp,20)
	fmt.Printf("表达式%s = %v", exp, res)
}
