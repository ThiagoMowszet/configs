package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	a := add(numbers)
	fmt.Println(a)
}

func add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum

}
