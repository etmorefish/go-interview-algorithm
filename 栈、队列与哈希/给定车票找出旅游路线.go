package main

import "fmt"

func main() {
	input := map[string]string{"西安": "成都", "北京": "上海", "大连": "西安", "上海": "大连"}
	PrintResult(input)
}

func PrintResult(input map[string]string) {
	// 反转input中的kv
	reverseInput := map[string]string{}
	for k, v := range input {
		reverseInput[v] = k
	}
	// 找到起点
	start := ""
	for k, _ := range input {
		if _, ok := reverseInput[k]; !ok {
			start = k
			break
		}
	}
	if start == "" {
		fmt.Println("input 不合理")
	} else {
		to := input[start]
		fmt.Print(start, "->", to)
		for {
			if _, ok := input[start]; ok {
				start = to
				to = input[start]
				if to == "" {
					break
				}
				fmt.Print(",", start, "->", to)

			} else {
				break
			}
		}
		fmt.Println()
	}
}
