package main

import "fmt"

/*
	在计算机科学中，分治法是一种很重要的算法。
	字面上的解释是分而治之，就是把一个复杂的问题分成两个或更多的相同或相似的子问题。
	直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。
	分治法一般使用递归来求问题的解。*/
func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(F(5))
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))
}

// F
/*1.如果每次递归都要对越来越长的链进行运算，那速度极慢，并且可能栈溢出，导致程序奔溃。*/
//2.如果每次递归都要对越来越长的链进行运算，那速度极慢，并且可能栈溢出，导致程序奔溃。 所以有另外一种写法，尾递归
//3.尾部递归是指递归函数在调用自身后直接传回其值，而不对其再加运算，效率将会极大的提高。
//4.当递归调用是整个函数体中最后执行的语句且它的返回值不属于表达式的一部分时，
//这个递归调用就是尾递归。尾递归函数的特点是在回归过程中不用做任何操作，这个特性很重要
//F
func F(n int) int {
	if n == 1 {
		return 1
	}
	return n * F(n-1)
}

//二分查找
func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid,low,high)
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

