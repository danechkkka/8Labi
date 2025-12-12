package main

import (
	"cmp"
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
