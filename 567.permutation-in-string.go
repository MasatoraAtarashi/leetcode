package main

func checkInclusion(s1 string, s2 string) bool {
	s1Cnt := make(map[string]int)
	for _, s := range s1 {
		s1Cnt[string(s)]++
	}

	start := 0
	end := len(s1) - 1
	for end <= len(s2)-1 {
		s2Cnt := make(map[string]int)
		for i := start; i <= end; i++ {
			s2Cnt[string(s2[i])]++
		}

		// 比較
		if compareMap(s1Cnt, s2Cnt) {
			return true
		}

		start++
		end++
	}

	return false
}

func compareMap(s1Cnt, s2Cnt map[string]int) bool {
	for key, value := range s1Cnt {
		if s2Cnt[key] != value {
			return false
		}
	}
	return true
}
