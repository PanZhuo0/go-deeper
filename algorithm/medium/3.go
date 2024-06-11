package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring(`bbbbb`))
}

func lengthOfLongestSubstring(s string) int {
	// sliding windows
	m := make(map[byte]bool) //judge repeated by map
	i := 0                   //slow
	j := 0                   //fast
	l := 0                   //length of sub string
	max_len := 0             //max length of all sub string
	for i < len(s) && j < len(s) {
		// try to reduce process times
		// if max length of sub string bigger than remind length of string
		if len(s)-i < max_len {
			break
		}
		// if not repeat
		if !m[s[j]] {
			m[s[j]] = true
			l++
			j++
		} else {
			// if repeat
			l--
			m[s[i]] = false
			i++
		}

		if l > max_len {
			max_len = l
		}
	}

	return max_len
}
