// Package mathutil は数値まわりの練習コードを置く場所。
package mathutil

// GCD は a, b の最大公約数を返す（ユークリッドの互除法）。負値は絶対値で扱う。
func GCD(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// FibSequence は f(0)=0, f(1)=1 から始まる先頭 n 個のフィボナッチ数列を返す。
// n <= 0 のときは空スライス。
func FibSequence(n int) []int64 {
	if n <= 0 {
		return []int64{}
	}
	out := make([]int64, 0, n)
	var a, b int64 = 0, 1
	for i := 0; i < n; i++ {
		out = append(out, a)
		a, b = b, a+b
	}
	return out
}
