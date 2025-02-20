package main

import (
	"fmt"
	"strconv"
	"strings"
)

const divider = "#"

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	// 文字列の数#文字数#...文字
	var ss string
	var lens string
	for _, s := range strs {
		lens += fmt.Sprintf("%d%s", len(s), divider)
		ss += s
	}

	return fmt.Sprintf("%d%s%s%s", len(strs), divider, lens, ss)
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	// まずは文字数把握
	splited := strings.Split(strs, divider)
	if len(splited) == 0 {
		return nil
	}

	strslen, err := strconv.Atoi(splited[0])
	if err != nil {
		panic(err)
	}

	slens := make([]int, 0, strslen)
	var allStrLen int
	for i := 1; i <= strslen; i++ {
		slen, err := strconv.Atoi(splited[i])
		if err != nil {
			panic(err)
		}
		slens = append(slens, slen)
		allStrLen += slen
	}

	// ↑を使って文字抽出
	var decoded []string
	allString := strs[len(strs)-allStrLen:]
	// fmt.Println(allString)
	for _, slen := range slens {
		fmt.Println("slen", slen)
		decoded = append(decoded, allString[:slen])
		allString = allString[slen:]
	}

	return decoded
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
