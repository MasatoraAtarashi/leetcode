package main

func characterReplacement(s string, k int) int {
	var start, end, max int
	for end <= len(s)-1 {
		maxFreq := 0
		freqMap := make(map[string]int)
		for i := start; i <= end; i++ {
			freqMap[string(s[i])]++
			if maxFreq < freqMap[string(s[i])] {
				maxFreq = freqMap[string(s[i])]
			}
		}

		windowLen := end - start + 1
		// kで置き換えられなかったら窓を左から縮める
		if windowLen-maxFreq > k {
			start++
		} else {
			// k回で置き換えられるなら、窓の幅が現状の最大。更新して窓を広げる
			if max < windowLen {
				max = windowLen
			}
			end++
		}

	}

	return max
}
