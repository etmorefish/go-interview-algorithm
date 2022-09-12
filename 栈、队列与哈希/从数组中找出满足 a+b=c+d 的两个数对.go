package main

import "fmt"

type Pair struct {
	first, second int
}

func FindPairs(arr []int) bool {
	sumPair := map[int]*Pair{}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if _, ok := sumPair[sum]; !ok {
				sumPair[sum] = &Pair{i, j}
			} else {
				p := sumPair[sum]
				fmt.Printf("(%v,%v), (%v,%v)",
					arr[i], arr[j],
					arr[p.first], arr[p.second])
				return true
			}
		}
	}
	return false
}

func main() {
	arr := []int{3, 4, 7, 8, 9, 10, 20}
	FindPairs(arr)
}
