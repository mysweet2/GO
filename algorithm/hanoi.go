package main

import "fmt"

var i int;

func main() {

	fmt.Println("阿飞哥做的汉诺塔函数")
	n := 3
	
	Hanoi(n, "A", "B", "C")
}

func move(n int, startPos string, endPos string) {
	i++
	fmt.Printf("第%v步:将%v号盘子%v---->%v\n", i, n, startPos, endPos)
}

// Hanoi 汉诺塔
func Hanoi(n int, startPos string, tranPos string, endPos string) {
	if n == 1 {
		move(n, startPos, endPos)
	} else {
		Hanoi(n-1, startPos, endPos, tranPos)
		move(n, startPos, endPos)
		Hanoi(n-1, tranPos, startPos, endPos)
	}
}
