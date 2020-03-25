package others

import "fmt"

func TypeJudge(items ...interface{}){
	for i,x := range items{
		switch  x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool 值是%v\n",i,x)
		case float64:
			fmt.Printf("param #%d is a float64 值是%v\n",i,x)
		case int, int64:
			fmt.Printf("param #%d is a int 值是%v\n",i,x)
		case nil:
			fmt.Printf("param #%d is a nil 值是%v\n",i,x)
		case string:
			fmt.Printf("param #%d is a string 值是%v\n",i,x)
		default:
			fmt.Printf("param #%d is unknown 值是%v\n",i,x)
		}
	}
}