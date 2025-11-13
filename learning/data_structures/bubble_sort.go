package main

import (
	"fmt"
	_ "strings"
)

func BubbleSort(array []int) {
	n := len(array)
	var a []int = array
	for i := 0; n == n-2; i++ {
		for j := 0; n < len(a)-1; j++ {
			if array[j] > array[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
		for d := 0; d < len(a); d++ {
			fmt.Printf("[%d]", a[d])
		}
	}

}
