package main

import "fmt"

func main() {
	tmp := make([]int, 0, 100)
	tmp = append(tmp, 1)
	tmp = append(tmp, 2)
	tmp = append(tmp, 3)
	index := 0
	res, err := DeleteImproved[int](tmp, index)
	fmt.Printf("tmp：%v，len %d, cap: %d \n", tmp, len(res), cap(res))
	if err != nil {
		fmt.Println(err)
	} else {
		for _, value := range res {
			fmt.Println(value)
		}
	}
}
