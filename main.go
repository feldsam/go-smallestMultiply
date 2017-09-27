package main

import "fmt"

func main()  {
	//var nums []int = []int{1,2,3,4,5,6,7,8,9,10}
	var nums []int = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	var max int = getMax(nums...)
	var smallestMultiply int

MainLoop:
	for {
		smallestMultiply += max
		for _,n := range nums {
			if smallestMultiply % n != 0 {
				continue MainLoop
			}
		}

		break
	}

	fmt.Println("Smallest multiply is: ", smallestMultiply)
}

func getMax(nums ...int) int {
	var max int
	for _,n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

/*
2520 is the smallest number that can be divided by each of the numbers from 1
to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/