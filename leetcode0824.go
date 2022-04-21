package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	var vowels = map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}
	ans := &strings.Builder{}
	for i, cnt, n := 0, 1, len(sentence); i < n; i++ {
		if cnt > 1 {
			ans.WriteByte(' ')
		}
		start := i
		for i++; i < n && sentence[i] != ' '; i++ {
		}
		cnt++
		if _, ok := vowels[sentence[start]]; ok {
			ans.WriteString(sentence[start:i])
		} else {
			ans.WriteString(sentence[start+1 : i])
			ans.WriteByte(sentence[start])
		}
		ans.WriteByte('m')
		ans.WriteString(strings.Repeat("a", cnt))
	}
	return ans.String()
}

func main() {
	str := "I speak Goat Latin"
	fmt.Println(toGoatLatin(str))
}
