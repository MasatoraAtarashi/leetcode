package main

func lengthOfLongestSubstring(s string) int {
	startIdx := 0
	seen := make(map[string]int)
	maxLength := 0

	for i, c := range s {
		// 今チェックしてるウィンドウにすでにある
		if idx, ok := seen[string(c)]; ok {
			if idx >= startIdx {
				startIdx = idx + 1
			}
		}

		seen[string(c)] = i
		length := i - startIdx + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
