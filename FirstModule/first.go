package main

import "FirstProject/FourtModule"

//  1
//  func IsEven(n int) bool {
//	  return n%2 == 0
//  }

//	2
//	func Max(a int, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}

//  3
//  func Factorial(n int) (fact int) {
//	  fact = 1
//
//	  for i := 1; n >= i; i++ {
//	  	fact *= i
//	  }
//	  return
//  }

//  4
//	func SumSlice(nums []int) (sum int) {
//		for _, value := range nums {
//			sum += value
//		}
//		return
//	}

//  5
//	func CountVowels(s string) (count int) {
//		vowelsArray := [5]rune{'a', 'e', 'i', 'o', 'u'}
//		count = 0
//
//		for _, wordLetter := range s {
//			for _, valueVowel := range vowelsArray {
//				if wordLetter == valueVowel {
//					count++
//				}
//			}
//		}
//
//		return
//	}

//  6
//  func IsPrime(n int) bool {
//	  for i := n - 1; i > 1; i-- {
//		  if n%i == 0 {
//		  	  return false
//		  }
//	  }
//	  return true
//  }

//  7
//  func Reverse(s string) (reverseString string) {
//
//	  for i := len(s) - 1; i >= 0; i-- {
//		  reverseString += string(s[i])
//	  }
//	  return
//
//  }

//  8
//  func Fibonacci(n int) int {
//	  prevNum, fibonacciNum := 0, 0
//
//	  for i := 0; i <= n; i++ {
//
//	  	  if i == 1 {
//			  fibonacciNum = 1
//		  } else {
//			  currentFibonacci := fibonacciNum
//			  fibonacciNum = fibonacciNum + prevNum
//			  prevNum = currentFibonacci
//		  }
//	  }
//	  return fibonacciNum
//  }

//  9
//  func FindMin(nums []int) (min int) {
//
//	  min = nums[0]
//
//	  for i := 0; i < len(nums); i++ {
//	  	  if nums[i] < min {
//		  	  min = nums[i]
//		  }
//	  }
//
//	  return min
//  }

// 10
//  func IsPalindrome(s string) bool {
//	  counterBack := 1
//
//	  for i := 0; i < len(s)/2; i++ {
//	  	  if s[i] != s[len(s)-counterBack] {
//			  return false
//		  }
//		  counterBack++
//	  }
//
//	  return true
//  }

func main() {
	FourtModule.Second()
}
