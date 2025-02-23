package main

func RemoveIndex[T any](slice []T, index int) []T {
	new := append([]T{}, slice...)
	return append(new[:index], new[index+1:]...)
}
