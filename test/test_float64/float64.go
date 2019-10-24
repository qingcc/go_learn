package test_float64

import (
	"unicode"
)

////四舍五入 取精度
//func ToFixed(f float64, places int) float64 {
//	shift := math.Pow(10, float64(places))
//	fv := 0.0000000001 + f //对浮点数产生.xxx999999999 计算不准进行处理
//	return math.Floor(fv*shift+.5) / shift
//}
//
//func RoundFloat(f float64, m int) float64 {
//	floatStr := fmt.Sprintf("%."+strconv.Itoa(m)+"f", f)
//	flt, _ := strconv.ParseFloat(floatStr, 64)
//	return flt
//}
//
//func BenchmarkTestFloat(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		RoundFloat(1.299999, 3)
//	}
//}

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
