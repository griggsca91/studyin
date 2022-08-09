package main

import (
	"fmt"

	"github.com/griggsca91/studyin/datastructure/stack"
)

func main() {
	fmt.Println("Enter size of queue")
	var size int
	fmt.Scanln(&size)

	stack := stack.NewArrayStack[int](size)

	for {
		fmt.Println("choose option:")
		fmt.Println("1. Add element")
		fmt.Println("2. Pop element")
		fmt.Println("3. Show stack")

		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			var value int
			fmt.Println("Enter integer to add")
			fmt.Scanln(&value)
			stack.Add(value)
		case 2:
			value, err := stack.Pop()
			if err != nil {
				fmt.Println("Stack empty", err)
			} else {
				fmt.Println("element popped:", value)
			}

		case 3:
			fmt.Println(stack)
		default:
			fmt.Println("Invalid input")
		}
	}
}
