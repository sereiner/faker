package base

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const Rex = `[\w]+`

const p = `#`

//RandomDigit 返回 0 到 9 之间随机数
func RandomDigit() int {
	return RandInt(0, 10)
}

//RandomDigitNotNull 返回 1 到 9 之间随机数
func RandomDigitNotNull() int {
	return RandInt(1, 10)
}

//RandomDigitNot 返回一个 0 到 9 之间的随机数,但是不包括 except
func RandomDigitNot(except int) int {
	res := NumberBetween(0, 9)
	if res >= except {
		res++
	}
	return res
}

//RandomFloat32 返回一个 float32
func RandomFloat32() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

//RandomFloat64 返回一个 float64
func RandomFloat64() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

//RandInt 返回 min 到 max 之间的随机数
func RandInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//NumberBetween .
func NumberBetween(min, max int) int {
	if max < min {
		min, max = max, min
	}
	return RandInt(min, max)
}

//RandomElements 从给定的数组中获取count个不重复的随机元素
func RandomElements(a []string, count int) (r []string) {
	if count == 0 || len(a) == 0 || count > len(a) {
		return nil
	}
	res := make([]string, len(a), len(a))
	ii := 0
	for {
		if ii == count {
			break
		}
		i := NumberBetween(0, len(a))
		s := a[i]
		if res[i] == "" {
			res[i] = s
			r = append(r, s)
			ii++
		}
	}
	return r
}

//RandomElement 从给定的数组中获取一个随机元素
func RandomElement(a []string) string {
	if len(a) == 0 {
		return ""
	}
	return RandomElements(a, 1)[0]
}

// Numerify 根据填充字符串,返回一个随机数,默认返回 5位
func Numerify(format ...string) string {
	re, _ := regexp.Compile(p)
	if len(format) > 0 && format[0] != "" {
		res := re.ReplaceAllStringFunc(format[0], func(s string) string {
			return strconv.Itoa(RandomDigit())
		})
		return res
	}
	return re.ReplaceAllStringFunc("#####", func(s string) string {
		return strconv.Itoa(RandomDigit())
	})

}

// Lexify 根据填充字符串,返回一个随机字符串,默认返回 5位
func Lexify(format ...string) string {
	re, _ := regexp.Compile(p)
	if len(format) > 0 && format[0] != "" {
		res := re.ReplaceAllStringFunc(format[0], func(s string) string {
			return string(RandInt(65, 122))
		})
		return res
	}
	return re.ReplaceAllStringFunc("#####", func(s string) string {
		return string(RandInt(65, 122))
	})
}
