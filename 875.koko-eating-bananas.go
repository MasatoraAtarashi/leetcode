package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	fmt.Println("pilis", piles)

	// まずは一番多い山出す
	maxp := 0
	for _, pile := range piles {
		if maxp < pile {
			maxp = pile
		}
	}
	fmt.Println("maxp: ", maxp)

	// 1~maxpの中で二分探索してh満たせる最小のkを探す
	mink := maxp
	i := 1
	j := maxp
	for i <= j {
		mid := (i + j) / 2
		fmt.Println("mid: ", mid)
		// 満たせるかチェック
		cnt := 0
		for _, pile := range piles {
			cnt += pile / mid
			if pile%mid != 0 {
				cnt += 1
			}
			if cnt > h {
				break
			}
		}
		fmt.Println("cnt: ", cnt)

		// 満たせててより小さかったらmink更新
		if cnt <= h {
			mink = mid
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	fmt.Println("mink: ", mink)

	return mink
}
