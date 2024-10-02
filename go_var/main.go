package main

import (
	"fmt"
	"math"
	"strings"
)

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func calhan() int {
	s := "hello沙河小王子"
	cnt := 0
	for _, r := range s {
		if r > 128 {
			cnt++
		}
	}
	return cnt

}
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
func print99() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
func sumsz() {
	var a = [5]int{1, 3, 5, 7, 8}
	sum := 0
	for i := range a {
		sum += a[i]
	}
	fmt.Println(sum)

}
func calzs() {
	s := "how do you do"
	v := strings.Split(s, " ")
	m := make(map[string]int)
	for i := range v {
		m[v[i]]++
	}
	for key, value := range m {
		fmt.Println(key, "=", value)
	}
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for i := range users {
		cnt := 0
		for j := 0; j < len(users[i]); j++ {
			if users[i][j] == 'e' || users[i][j] == 'E' {
				coins -= 1
				cnt += 1
			}
			if users[i][j] == 'i' || users[i][j] == 'I' {
				coins -= 2
				cnt += 2
			}
			if users[i][j] == 'o' || users[i][j] == 'O' {
				coins -= 3
				cnt += 3
			}
			if users[i][j] == 'u' || users[i][j] == 'U' {
				coins -= 4
				cnt += 4
			}

		}
		distribution[users[i]] = cnt
		fmt.Println(users[i], cnt)
	}
	return coins
}

type stu struct {
	id    int
	name  string
	age   int
	score int
}
type stuList struct {
	students []stu
}

func (s *stuList) addStu(a stu) {
	s.students = append(s.students, a)
}

func (s *stuList) changeStu(a stu) {
	for index, v := range s.students {
		if v.id == a.id {

			s.students[index] = a
		}
	}
}
func (s *stuList) deleteStu(a stu) {
	for index, v := range s.students {
		if v.id == a.id {

			s.students = append(s.students[:index], s.students[index+1:]...)
		}
	}
}
func (s *stuList) showStu() {
	for _, v := range s.students {

		fmt.Printf("%d %s %d %d\n", v.id, v.name, v.age, v.score)
	}
}
func main() {
	s := stuList{}
	st := stu{1, "s", 2, 1}
	s.addStu(st)
	st = stu{2, "sd", 2, 1}
	s.addStu(st)
	s.showStu()
	fmt.Printf("%#v %#v\n", st, s)
	st = stu{1, "sgfdrg", 2, 1}
	s.changeStu(st)
	s.deleteStu(st)
	fmt.Printf("%#v %#v\n", st, s)
}
