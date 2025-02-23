package main

var digitsLetterMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var result []string
	var current string
	var fff func(ds string)
	fff = func(ds string) {
		if len(ds) == 0 {
			result = append(result, current)
			return
		}

		for _, l := range digitsLetterMap[string(ds[0])] {
			current += l
			fff(ds[1:])
			current = current[:len(current)-1]
		}
	}

	fff(digits)
	return result
}
