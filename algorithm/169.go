package main

func main() {

}

func majorityElement(nums []int) int {
	// Boyer-moore 算法
	// candidate
	count := 0
	var candidater int
	for _, n := range nums {
		if count == 0 {
			// if count==0 we should change candidater
			candidater = n
		}
		if n == candidater {
			// if vote up candidate
			count++
		} else {
			// not vote up candidate equivalent to  vote no
			// not vote up candidate means vote other person
			count--
		}
		// eventually,we will get the majority element
	}
	return candidater
}
