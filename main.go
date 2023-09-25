package main

import (
	"data/datastructures/stack"
	"fmt"
)

func main() {
	nums := stack.Stack{}

	for i := 0; i < 10; i++ {
		nums.Push(i + 1)
	}

	for _, num := range nums {
		fmt.Printf("%T, %d\n", num, num)
	}

	fmt.Println(nums)

	names := stack.Stack{}

	names.Push("Robert")
	names.Push("Fred")
	names.Push("Alicia")
	fmt.Println(names)

	for _, name := range names {
		fmt.Printf("%T, %s\n", name, name)
	}

	names.Pop()

	fmt.Println(names)
}
