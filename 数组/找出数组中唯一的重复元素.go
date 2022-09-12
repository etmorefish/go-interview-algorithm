package main

import "fmt"

func FindDupByHash(arr []int) int {
	if arr == nil {
		return -1
	}
	data := map[int]bool{}
	for _, v := range arr {
		if _, ok := data[v]; ok {
			return v
		} else {
			data[v] = true
		}
	}
	return -1
}

func FindDupByXOR(arr []int) int {
	if arr == nil {
		return -1
	}
	result := 0
	len := len(arr)
	for _, v := range arr {
		result ^= v
	}
	for i := 1; i < len; i++ {
		result ^= i
	}
	return result
}

func FindDupByLoop(arr []int) int {
	if arr == nil {
		return -1
	}
	slow := 0
	fast := 0

	for slow != fast {
		fast = arr[arr[fast]] // fast一次走两步
		slow = arr[slow]      // slow一次走1步
	}
	fast = 0
	for slow != fast {
		fast = arr[fast]
		slow = arr[slow]
	}
	return slow
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 5}
	// fmt.Println(FindDupByHash(arr))
	// fmt.Println(FindDupByXOR(arr))
	fmt.Println(FindDupByLoop(arr))

}
