package main

import (
	"fmt"
	"strings"
)

func isAnagram(s string, t string) bool {
	for _, ss := range s {
		if !strings.Contains(t, string(ss)) {
			return false
		}

		t = strings.Replace(t, string(ss), "", 1)
		fmt.Println(t)
	}

	if len(t) == 0 {
		return true
	}
	return false
}
