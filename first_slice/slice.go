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
	// 子切片不发生扩容的情况下，和原切片的容量相等
	return Shrink(append(sli[:index], sli[index+1:]...)), nil
}

// Shrink 这是缩容
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

func calCapacity(c, l int) (int, bool) {
	// 容量 <=64 缩不缩都无所谓，因为浪费内存也浪费不了多少
	// 你可以考虑调大这个阈值，或者调小这个阈值
	if c <= 64 {
		println("yse")
		return c, false
	}
	// 如果容量大于 2048，但是元素不足一半，
	// 降低为 0.625，也就是 5/8
	// 也就是比一半多一点，和正向扩容的 1.25 倍相呼应
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	// 如果在 2048 以内，并且元素不足 1/4，那么直接缩减为一半
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	// 整个实现的核心是希望在后续少触发扩容的前提下，一次性释放尽可能多的内存
	return c, false
}
