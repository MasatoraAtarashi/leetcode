package main

func RemoveIndex[T any](slice []T, index int) []T {
	new := append([]T{}, slice...)
	return append(new[:index], new[index+1:]...)
}

func SliceContains[T comparable](slice []T, value T) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}
