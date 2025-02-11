package main

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	var start, max, maxFreq int
	for end := 0; end < len(s); end++ {
		freq[s[end]-'A']++
		if freq[s[end]-'A'] > maxFreq {
			maxFreq = freq[s[end]-'A']
		}
		windowLen := end - start + 1
		// kで置き換えられなかったら窓を左から縮める
		if windowLen-maxFreq > k {
			freq[s[start]-'A']--
			start++
			continue
		}

		// k回で置き換えられるなら、窓の幅が現状の最大。更新して窓を広げる
		if max < windowLen {
			max = windowLen
		}
	}

	return max
}
