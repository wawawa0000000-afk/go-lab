// Package stats - 07: テストの題材となる小さなライブラリ。
// テスト実行: go test ./reference/07-testing
package stats

import (
	"errors"
	"sort"
	"strings"
)

var ErrEmpty = errors.New("stats: 空の入力")

// Mean は平均を返す。空なら ErrEmpty。
func Mean(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrEmpty
	}
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums)), nil
}

// Median は中央値を返す。入力は変更しない。
func Median(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrEmpty
	}
	cp := append([]float64(nil), nums...) // 防御的コピー
	sort.Float64s(cp)
	mid := len(cp) / 2
	if len(cp)%2 == 1 {
		return cp[mid], nil
	}
	return (cp[mid-1] + cp[mid]) / 2, nil
}

// WordFrequency は文章中の単語出現回数を返す（小文字化・記号除去）。
func WordFrequency(text string) map[string]int {
	freq := map[string]int{}
	for _, w := range strings.Fields(text) {
		w = strings.ToLower(strings.Trim(w, ".,!?\"'()[]:;"))
		if w != "" {
			freq[w]++
		}
	}
	return freq
}

// IsPalindrome は大文字小文字と空白を無視して回文か判定する。
func IsPalindrome(s string) bool {
	var r []rune
	for _, c := range strings.ToLower(s) {
		if c != ' ' {
			r = append(r, c)
		}
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}
