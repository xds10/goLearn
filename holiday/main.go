// holiday/main.go

package main

import (
	"fmt"
	"holiday/summer"

	"github.com/q1mi/hello"
)

func main() {
	fmt.Println("现在是假期时间...")
	hello.SayHi()
	a := summer.Add(1, 2)
	fmt.Println(a)
}
