package main

func main() {
	fmt.Println(arrangeCoins(10))
}

func arrangeCoins(n int) int {
	for i := 1; i < n; i++ {
		n = n - i
		if n <= i {
			return i
		}
	}
	return n
}
