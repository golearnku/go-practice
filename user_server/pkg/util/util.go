package util

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golearnku/go-practice/listen1/pkg/util"
)

func CreateCaptcha() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

// GetCilentIp 返回远程客户端的 IP，如 192.168.1.1
func GetCilentIp(r *http.Request) (ip string) {
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return
	}
	ip, _, _ = net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	return
}

// 一天中的最后一秒钟
func LastSecondOfDay(day time.Time) time.Time {
	local, _ := time.LoadLocation("Local")
	return time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, local)
}

// GetUsername 获取用户名
func GetUsername() string {
	// fixme: redis INCR 获取 获取自增 ID
	var userID int64
	idxStr := util.DecimalToAny(util.PseudoEncrypt(userID), 64)
	for i, l := 0, 8-len(idxStr); i < l; i++ {
		idxStr = "u" + idxStr
	}
	return "user_" + idxStr
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
