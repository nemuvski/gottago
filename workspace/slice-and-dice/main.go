package main

import (
	"fmt"
)

func sliceAndDice(src []int, size uint) [][]int {
	if size == 0 || len(src) == 0 {
		return [][]int{}
	}

	_blocks := make([][]int, 0)
	
	for int(size) < len(src) {
		src, _blocks = src[size:], append(_blocks, src[0:size:size])
	}
	// 余った分を追加
	_blocks = append(_blocks, src)

	return _blocks
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	blocks := sliceAndDice(arr, 3)
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(blocks, len(blocks), cap(blocks))
}
