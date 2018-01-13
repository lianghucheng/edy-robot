package common

import (
	math "math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	Chars = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c",
		"d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C",
		"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "~", "!", "@",
		"#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[",
		"]", "{", "}", "|", "<", ">", "?", "/", ".", ",", ";", ":"}

	NumberChars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Shuffle(a []string) []string {
	n := len(a)
	if n == 0 {
		return a
	}
	b := make([]string, n)

	m := rand.Perm(n)
	for i := 0; i < n; i++ {
		b[i] = a[m[i]]
	}
	return b
}

func SplitHeadimgurl(headimgurl string) []string {
	var imgurl []string
	url := headimgurl[:41]
	number := headimgurl[41:42]
	jpg := headimgurl[42:]
	for i := 0; i < 100; i++ {
		number = strconv.Itoa(i)
		imgurl = append(imgurl, url+number+jpg)
	}
	return imgurl
}

func Shuffle2(a []string) []string {
	i := len(a) - 1
	for i > 0 {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
		i--
	}
	return a
}

func GetToken(n int) string {
	if n < 1 {
		return ""
	}
	var tokens []string
	for i := 0; i < n; i++ {
		tokens = append(tokens, Chars[rand.Intn(90)]) // 90 是 Chars 的长度
	}
	return strings.Join(tokens, "")
}

// id 的第一位从 1 开始
func GetID(n int) int {
	if n < 1 {
		return -1
	}
	min := math.Pow10(n - 1)
	id := int(min) + rand.Intn(int(math.Pow10(n)-min))
	return id
}

func Index(a []int, sep int) int {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == sep {
			return i
		}
	}
	return -1
}

// 判断 value 是否在 array 中
func InArray(a []int, sep int) bool {
	for _, v := range a {
		if sep == v {
			return true
		}
	}
	return false
}

// 从 array 中移除最开始出现的 value
func RemoveOnce(a []int, sep int) []int {
	i := Index(a, sep)
	if i == -1 {
		return a
	}
	b := []int{}
	if i == 0 {
		b = append(b, a[1:]...)
	} else if i == len(a)-1 {
		b = append(b, a[:i]...)
	} else {
		b = append(b, a[:i]...)
		b = append(b, a[i+1:]...)
	}
	return b
}

func Remove(a []int, sub []int) []int {
	for _, v := range sub {
		a = RemoveOnce(a, v)
	}
	return a
}

func ReplaceAll(a []int, old, new int) []int {
	if old == new {
		return a
	}
	if InArray(a, old) {
		b := []int{}
		for _, v := range a {
			if old == v {
				b = append(b, new)
			} else {
				b = append(b, v)
			}
		}
		return b
	}
	return a
}

func Deduplicate(a []int) []int {
	n := len(a)
	if n == 0 {
		return a
	}
	m := make(map[int]bool)

	b := []int{a[0]}
	m[a[0]] = true
	for i := 1; i < n; i++ {
		if !m[a[i]] {
			b = append(b, a[i])
			m[a[i]] = true
		}
	}
	return b
}

// 比较两个数组的元素是否相等
func Equal(x, y []int) bool {
	if len(x) == len(y) {
		return Contain(x, y)
	}
	return false
}

func Contain(x, y []int) bool {
	if len(x) < len(y) {
		return false
	}
	temp := Deduplicate(y)
	for _, v := range temp {
		if Count(x, v) < Count(y, v) {
			return false
		}
	}
	return true
}

func Count(a []int, sep int) int {
	count := 0
	for _, v := range a {
		if sep == v {
			count++
		}
	}
	return count
}
