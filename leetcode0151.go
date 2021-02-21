package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	s = preProcess(s)
	wordarry := strings.Split(s, " ")
	i := 0
	l := len(wordarry) - 1
	for i < l {
		wordarry[i], wordarry[l] = wordarry[l], wordarry[i]
		i++
		l--
	}
	res := strings.Join(wordarry, " ")
	return res
}

func preProcess(s string) string {
	l := len(s)
	var res []byte
	flag := 1 //空格处理
	for l != 0 && s[l-1] == ' ' {
		l--
	}
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		if s[i] == ' ' && flag == 0 {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}

func main() {
	s := "a good   example"
	fmt.Println(s)
	s1 := "  Bob    Loves  Alice   "
	fmt.Println(s1)
}
