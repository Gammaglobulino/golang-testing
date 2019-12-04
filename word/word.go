//provides functions for working with string words
package word

import (
	"strings"
)

//counting the words on a string and return the usage of every single word
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//count the number of words contained in a string
func Count(s string) int {
	return len(strings.Fields(s))
}
