package base

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const Rex = `[\w]+`

const p = `#`

//RandomDigit return a random number between 0 and 9
func RandomDigit() int {
	return RandInt(0, 10)
}

//RandomDigitNotNull return a random number between 1 and 9
func RandomDigitNotNull() int {
	return RandInt(1, 10)
}

//RandomDigitNot .
func RandomDigitNot(except int) int {
	res := NumberBetween(0, 9)
	if res >= except {
		res++
	}
	return res
}

//RandomFloat32 .
func RandomFloat32() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

//RandomFloat64 .
func RandomFloat64() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

//RandInt .
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

func Numerify(format string) string{
	re, _ := regexp.Compile(p)
	res := re.ReplaceAllStringFunc(format, func(s string) string {
		return strconv.Itoa(RandomDigit())
	})
	return res
}
