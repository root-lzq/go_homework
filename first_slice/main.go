package main

import "fmt"

func main() {
	tmp := []int{1, 2, 3, 4, 5}
	index := 3

	res, err := DeleteImproved[int](tmp, index)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, value := range res {
			fmt.Println(value)
		}
	}
}
