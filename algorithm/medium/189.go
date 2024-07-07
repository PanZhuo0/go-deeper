package main

func main() {
	d := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(d, 3)
}

/* 先翻前len()-k个，再翻后k个，哎，整体翻一遍，哎你猜怎么着，哎哟喂，既然过了,这您受得了吗? */
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	l := len(nums)

	k = k % l
	// 前n-k个
	for i := 0; i < (l-k)/2; i++ {
		// i=0 , l-k-i=7-3-0
		nums[i], nums[l-k-1-i] = nums[l-k-1-i], nums[i]
	}
	// 后k个
	// k=3,l=7
	for i := l - k; i < (l-k+l)/2; i++ {
		nums[i], nums[2*l-k-i-1] = nums[2*l-k-i-1], nums[i]
	}

	// 全部翻转
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
	/* 关于部分反转的逻辑
	m->n反转
	for i:=m;i<(m+n)/2;i++{
		nums[i],nums[2m+n-i]=nums[2m+n-i],nums[i]
	}

	总之:
		1.i:= 反转的起始index
		2.遍历终止条件为: 小于起始(index+结束index)/2
		3.替换的两个index为
			nums[i] = nums[xxx-i]  (只要当i=起始index时,xxx-i的整体值等于结束index-1即可)
	*/
}
