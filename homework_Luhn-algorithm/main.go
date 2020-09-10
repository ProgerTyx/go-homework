package main

import (
	"fmt"
)

func reverse(sl []int) []int {
	var res []int
	for i := len(sl) - 1; i >= 0; i-- {
		res = append(res, sl[i])
	}
	return res
}

func sum(sl []int) int {
	var res int
	for _, v := range sl {
		res = res + v
	}
	return res
}

func luhnAlg(sl []int) bool {
	revSl := reverse(sl)
	for i := 1; i <= len(revSl); i++ {
		if i%2 != 0 {
			val := revSl[i] * 2
			if val > 9 {
				revSl[i] = val - 9
			} else {
				revSl[i] = val
			}
		}
	}
	if sum(revSl)%10 == 0 {
		return true
	}
	return false
}

func main() {
	var sl1 = []int{4, 4, 4, 1, 1, 1, 4, 4, 2, 2, 6, 0, 8, 1, 9, 8}
	var sl2 = []int{4, 5, 6, 1, 2, 6, 1, 2, 1, 2, 3, 4, 5, 4, 6, 7}
	var sl3 = []int{4, 5, 6, 1, 2, 6, 1, 2, 1, 2, 3, 4, 5, 4, 6, 4}
	fmt.Println(luhnAlg(sl1))
	fmt.Println(luhnAlg(sl2))
	fmt.Println(luhnAlg(sl3))
}
