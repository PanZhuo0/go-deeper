package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic(`\\u0000\\u0000`, `\\u0000c`))
}

// 该题目中实际上要求两个字符串彼此双射（一对一映射）
// 由此，对于两个字符串，对每个字符串维护一个映射表(K,V),
// 如果发现映射不符合则return false，整个字符串都通过映射检查才算通过,return true
func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte, 0)
	t2s := make(map[byte]byte, 0)

	for i := 0; i < len(t); i++ {
		si := s[i]
		ti := t[i]
		// 如果已有该char的映射关系，并且现在检查的映射关系与之前保存的映射关系不匹配
		if s2t[si] > 0 && s2t[si] != ti || t2s[ti] > 0 && t2s[ti] != si {
			return false
		}
		// 更新该char的映射关系
		s2t[si] = ti
		t2s[ti] = si
	}
	return true
}
