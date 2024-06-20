package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3749))
}

// 朴素法
func intToRoman(num int) string {
	// symbol := []string{"M", `CM`, `D`, `CD`, `C`, `XC`, `L`, `XL`, `X`, `IX`, `V`, `IV`, `I`}
	ret := []string{}
	for num > 0 {
		if num > 1000 {
			num -= 1000
			fmt.Println(num)
			ret = append(ret, `M`)
			continue
		} else if num > 900 {
			num -= 900
			fmt.Println(num)
			ret = append(ret, `CM`)
			continue
		} else if num > 500 {
			num -= 500
			fmt.Println(num)
			ret = append(ret, `D`)
			continue
		} else if num > 400 {
			num -= 400
			fmt.Println(num)
			ret = append(ret, `CD`)
			continue
		} else if num > 100 {
			num -= 100
			fmt.Println(num)
			ret = append(ret, `C`)
			continue
		} else if num > 90 {
			num -= 90
			fmt.Println(num)
			ret = append(ret, `XC`)
			continue
		} else if num > 50 {
			num -= 50
			fmt.Println(num)
			ret = append(ret, `L`)
			continue
		} else if num > 40 {
			num -= 40
			fmt.Println(num)
			ret = append(ret, `XL`)
			continue
		} else if num > 10 {
			num -= 10
			fmt.Println(num)
			ret = append(ret, `X`)
			continue
		} else if num > 9 {
			num -= 9
			fmt.Println(num)
			ret = append(ret, `IX`)
			continue
		} else if num > 5 {
			num -= 5
			fmt.Println(num)
			ret = append(ret, `V`)
			continue
		} else if num > 4 {
			num -= 4
			fmt.Println(num)
			ret = append(ret, `IV`)
			continue
		} else if num >= 1 {
			num -= 1
			fmt.Println(num)
			ret = append(ret, `I`)
			continue
		}
	}
	r := ``
	for i := 0; i < len(ret); i++ {
		r += ret[i]
	}
	return r
}

/*
模拟法

M：1000	CM：900
D：500	CD：400
C：100	XC：90
L：50	XL:40
X:10	 IX:9
V:5		 IV:4
I:1

每次都尽可能选择较大的数
*/
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman2(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}

	return string(roman)
}

/*
硬编码法
由于
	罗马数字千位，只能用M表示
	百位只能用C,CD,D,CM表示
	十位只能用X,XL,L,XC表示
	个位数字只能由I，IV,V,IX表示
因此
	对于num中的各个位，编码的结果互不影响，可以建立一个硬编码表，进行编码
	只需要查询表格+拼凑即可
*/

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman3(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
