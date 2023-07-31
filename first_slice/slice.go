package main

import (
	"errors"
)

/*
删除切片特定下标元素
时间复杂度 O(N); 空间复杂度O(1)
*/
func Delete[T any](sli []T, index int) ([]T, error) {
	if index < 0 || index >= len(sli) {
		return nil, errors.New("attempt to delete incorrect slice index")
	}
	if len(sli) == 1 {
		return []T{}, nil
	}
	res := make([]T, len(sli)-1)
	for i := 0; i < len(res); i++ {
		if i < index {
			res[i] = sli[i]
		} else {
			res[i] = sli[i+1]
		}
	}
	return res, nil
}

/*
	删除切片特定下标元素(改进版)
	时间复杂度 O(N); 空间复杂度O(1)
*/

func DeleteImproved[T any](sli []T, index int) ([]T, error) {
	if index < 0 || index >= len(sli) {
		return nil, errors.New("attempt to delete incorrect slice index")
	}
	return append(sli[:index], sli[index+1:]...), nil
}
