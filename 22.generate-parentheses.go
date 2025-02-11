package main

func generateParenthesis(n int) []string {
	var result []string
	var backtracking func(open, close int, str string)
	backtracking = func(open, close int, str string) {
		// fmt.Println(open, close, str)

		// 全部使い切ったら追加
		if open == n && close == n {
			result = append(result, str)
			return
		}

		// (を追加可能な条件→（がまだある
		if open <= n {
			backtracking(open+1, close, str+"(")
		}

		// )を追加可能な条件→(の方が多く使われている
		if open > close {
			backtracking(open, close+1, str+")")
		}
	}

	backtracking(0, 0, "")
	return result
}
