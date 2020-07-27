* 现有整型数组 a 、整型数组 b、以及整型v。请使用指定编程语言编写函数，判 断是否可以从 a 中选择一个数，b 中选择一个数，二者相加等于 v，如可以返回
  true ，否则返回 false 。比如
  
```
a = [10, 40, 5, 280];
b = [234, 5, 2, 148, 23]; v = 42;
```

将返回 true，因为a中40和b中2相加为42。

```go
package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 40, 5, 280}
	nums2 := []int{234, 5, 2, 148, 23}
	target := 42
	fmt.Println(twoSum2(nums, nums2, target))
}

func twoSum2(nums, nums2 []int, target int) bool {
	arr := make([]int, 0, len(nums)+len(nums2))
	arr = append(arr, nums...)
	arr = append(arr, nums2...)
	m := map[int]int{}
	for _, val := range arr {
		if _, ok := m[target-val]; ok {
			return true
		}
		m[val] = val
	}
	return false
}
```


* 现要求游戏玩家的用户名是 `user` 后加上不重复且随机的正整数，请参考 PostgreSQL 中 `pseudo_encrypt()` 的原理，使用指定编程语言，实现一个高效的用户名生成算法。

```go
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for i := 1000; i <= 9999; i++ {
		fmt.Println("user_" + strconv.FormatInt(PseudoEncrypt(int64(i)), 10))
	}
}

// PseudoEncrypt
func PseudoEncrypt(value int64) int64 {
	var (
		l1, l2, r1, r2 int64
	)

	l1 = (value >> 16) & 0xffff
	r1 = value & 0xffff
	for i := 0; i < 3; i++ {
		l2 = r1
		r := math.Round(float64((1366 * r1 + 150889) % 714025) / 714025.0 * 32767)
		r2 = l1 ^ int64(r)
		l1 = l2
		r1 = r2
	}
	return (r1 << 16) + l1
}
```

* 如果要求用户名均为8位，并且即便知道所用算法，也无法预测用户名序列，应该如何改进算法?

答：

```go


package main

import (
    "fmt"
	"math"
	"strconv"
)

func main() {
	for i := 10000000; i <= 20000000; i++ {
		idxStr := DecimalToAny(PseudoEncrypt(int64(i)), 64)
		for i, l := 0, 8-len(idxStr); i < l; i++ {
			idxStr = "u" + idxStr
		}
		userName := "user_" + idxStr
		fmt.Printf("随机生成的用户名: %s\n", userName)
	}
}

// PseudoEncrypt
// 参考 Pseudo encrypt 算法：https://wiki.postgresql.org/wiki/Pseudo_encrypt
func PseudoEncrypt(value int64) int64 {
	var (
		l1, l2, r1, r2 int64
	)

	l1 = (value >> 16) & 65535
	r1 = value & 65535
	for i := 0; i < 3; i++ {
		l2 = r1
		r := math.Round(float64((1366*r1+150889)%714025) / 714025.0 * 32767)
		r2 = l1 ^ int64(r)
		l1 = l2
		r1 = r2
	}
	return (r1 << 16) + l1
}

// DecimalToAny 10 进制转换为任意进制
func DecimalToAny(num int64, n int) string {
	var tenToAny = map[int64]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: "A", 37: "B", 38: "C", 39: "D", 40: "E", 41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N", 50: "O", 51: "P", 52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "V", 58: "W", 59: "X", 60: "Y", 61: "Z",
	}
	newNumStr := ""
	var remainder int64
	var remainderString string
	for num != 0 {
		remainder = num % int64(n)
		if 62 > remainder && remainder > 9 {
			remainderString = tenToAny[remainder]
		} else {
			remainderString = strconv.FormatInt(remainder, 10)
		}
		newNumStr = remainderString + newNumStr
		num = num / int64(n)
	}
	return newNumStr
}
```

   > 想到另外一种方案，使用预生成的方式，基于 redis 的 list 存储，每次获取的时候 pop 一个出来就行了.

* 请设计系统和协议，使用以上算法，实现手机用户的注册和登录，要求能够实现 用户名与设备的绑定，并应对 10万用户/小时 的注册和登录压力，系统应具有横 向扩展能力。

答： 首先没太理解是 设置的系统要每小时能支撑10万用户注册，还是峰值要能支撑10万用户注册，如果是要求每小时能支撑10万用户注册，那么平均到每秒注册用户就是：100000 / 3600 约等于 1 秒有 28 个用户注册；
所以对系统的要求不是很大，一般的系统都能支持这个体量。如果是峰值 10万用户注册，可能这 10w 用户发生在某一段时间内，这时候
需要给服务增加限流机制，避免整个服务不可用。

* 请改进系统设计，允许用户更换手机时保留账号。

    > 设计的注册和登录请查看 `user_server` [目录](https://github.com/golearnku/go-practice/tree/master/user_server)

