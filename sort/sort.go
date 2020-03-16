// 按照某个字段对结构体数组进行排序例子
package sort

import (
	"fmt"
	"math/rand"
	"sort"
)

type student struct {
	name string
	age int
	score int
}
type stu []student

func (s stu) Len() int {
	return len(s)
}

func (s stu) Less(i,j int) bool {
	return s[i].score < s[j].score
}

func (s stu) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func test() {
	var sts stu
	for i := 0; i < 10; i++{
		st :=student{
			name:  fmt.Sprintf("学生~%d",rand.Intn(100)),
			age:   rand.Intn(50),
			score: rand.Intn(100),
		}
		sts = append(sts,st)
	}
	for _,v := range sts {
		fmt.Println(v)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	sort.Sort(sts)
	for _,v := range sts {
		fmt.Println(v)
	}

}

/*
{学生~81 37 47}
{学生~59 31 18}
{学生~25 40 56}
{学生~0 44 11}
{学生~62 39 28}
{学生~74 11 45}
{学生~37 6 95}
{学生~66 28 58}
{学生~47 47 87}
{学生~88 40 15}
~~~~~~~~~~~~~~~~~~~~~~~~~~
{学生~0 44 11}
{学生~88 40 15}
{学生~59 31 18}
{学生~62 39 28}
{学生~74 11 45}
{学生~81 37 47}
{学生~25 40 56}
{学生~66 28 58}
{学生~47 47 87}
{学生~37 6 95}
*/