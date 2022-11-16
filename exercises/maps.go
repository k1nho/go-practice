package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	v := strings.Fields(s);
	
	for _, w := range v {
		val, ok := m[w]
		if(ok) {
			m[w] = val +1;
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
