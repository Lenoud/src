package main

import "fmt"

func main() {
	//定义一个整形
	var age int = 18
	fmt.Printf("%T,%d\n", age, age)
	//定义应该浮点型
	var money float64 = 3.19
	fmt.Printf("%T,%.1f\n", money, money)
}
