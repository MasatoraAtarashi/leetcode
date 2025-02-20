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
	// fmt.Println(strs)

	var decoded []string
	splited := strings.Split(strs, divider)
	if len(splited) == 0 {
		return nil
	}

	strslen, err := strconv.Atoi(splited[0])
	if err != nil {
		panic(err)
	}
	// fmt.Println(strslen)

	ss := splited[len(splited)-1]
	for i := 1; i <= strslen; i++ {
		// fmt.Println(i)
		slen, err := strconv.Atoi(splited[i])
		if err != nil {
			panic(err)
		}

		if slen == 0 {
			decoded = append(decoded, "")
			continue
		}
		if ss == "" {
			decoded = append(decoded, divider)
		} else {
			decoded = append(decoded, ss[:slen])
			ss = ss[slen:]
		}
		// fmt.Println("decoded", decoded)
		// fmt.Println("ss", ss)
	}

	return decoded
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
