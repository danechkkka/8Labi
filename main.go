package main

import (
	"cmp"
	"fmt"
	"sync"
)

func Max[T cmp.Ordered](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice)-1; j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {
	slice := make([]int, 0, 15)
	n_slice := make([]int, 3)
	parts := 5
	partSize := len(slice) / parts
	var wg sync.WaitGroup
	for i := 0; i < len(slice); i += partSize {
		end := i + partSize
		if i == len(slice) {
			end = len(slice) - 1
		}
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			s = Max(s)
			n_slice = append(n_slice, s[0])
		}(slice[i:end])
	}
	wg.Wait()
	Max(n_slice)
	fmt.Print(n_slice[0])
}
