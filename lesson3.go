package main

import "fmt"

func main(){
	fmt.Println("true && false", true && false)
	fmt.Println("")

	i:=1
	for i <= 10{
		fmt.Println(i)
		i++
	}

	for j := 0;j<5;j++{
		fmt.Println(j)
	}
}