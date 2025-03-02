package main

func wordBreak(s string, wordDict []string) bool {
	breakable := make([]bool, len(s)+1)
	// 空文字は常に分割可能
	breakable[0] = true

	for i := range s {
		// fmt.Println("i", i)
		// fmt.Println("check", string(s[:i+1]))
		// fmt.Println("breakable", breakable)
		for j := 0; j < i; j++ {
			// fmt.Println("j", j)
			// fmt.Println("前半", string(s[:j]))
			// fmt.Println("後半", string(s[j:i+1]))
			// fmt.Println("")
			if breakable[j] && SliceContains(wordDict, string(s[j:i+1])) {
				// fmt.Printf("target: %s is breakable\n", string(s[:i+1]))
				breakable[i+1] = true
				break
			}
		}
	}

	return breakable[len(s)]
}
