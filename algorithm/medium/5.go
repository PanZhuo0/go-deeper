package main

/*
	longestPalindrome	寻找字符串s的最大回文子串
	--枚举法--
	paramter:
		s string 需要检查的字符串s

	response:
		res string:最大回文子串
*/
func longestPalindrome(s string) string {
	res := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > len(res) && isPalindrome(s, i, j) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

/*
	isPalindrome 用于判断字符串s[left:right+1]这个范围是否为回文串
	是返回		true
	不是返回	false
*/
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
	动态规划
*/
func longestPalindrome2(s string) string {

}

/*
longestPalindrome 最长回文子串
Algorithm:中心扩散法
*/
func longestPalindrome3(s string) string {
	res := s[0:1]
	for i := 0; i < len(s); i++ {
		r1 := extendFromCenter(s, i, i)
		r2 := extendFromCenter(s, i, i+1)
		if r1 > res {
			res = r1
		} else if r2 > res {
			res = r2
		}
	}
	return res
}

func extendFromCenter(s string, left, right int) string {
	for left >= 0 && right <= len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left : right+1]
}
