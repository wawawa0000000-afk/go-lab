package stats

import (
	"errors"
	"fmt"
	"maps"
	"testing"
)

// --- 1. 基本的なテスト ---
func TestMean(t *testing.T) {
	got, err := Mean([]float64{1, 2, 3, 4})
	if err != nil {
		t.Fatalf("予期しないエラー: %v", err)
	}
	if got != 2.5 {
		t.Errorf("Mean = %v, want 2.5", got)
	}
}

// --- 2. エラーケース ---
func TestMean_Empty(t *testing.T) {
	_, err := Mean(nil)
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}

// --- 3. テーブル駆動テスト（Go の定番スタイル）---
func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		in   []float64
		want float64
	}{
		{"奇数個", []float64{3, 1, 2}, 2},
		{"偶数個", []float64{4, 1, 3, 2}, 2.5},
		{"1個", []float64{7}, 7},
		{"重複あり", []float64{5, 5, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Median(tt.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Median(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

// --- 4. 入力を破壊していないことの確認 ---
func TestMedian_DoesNotMutate(t *testing.T) {
	in := []float64{3, 1, 2}
	_, _ = Median(in)
	if in[0] != 3 || in[1] != 1 || in[2] != 2 {
		t.Errorf("入力が書き換えられた: %v", in)
	}
}

// --- 5. map の比較 ---
func TestWordFrequency(t *testing.T) {
	got := WordFrequency("The cat, the CAT! the dog.")
	want := map[string]int{"the": 3, "cat": 2, "dog": 1}
	if !maps.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"level":                       true,
		"A man a plan a canal Panama": true,
		"golang":                      false,
		"":                            true,
	}
	for in, want := range cases {
		if got := IsPalindrome(in); got != want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}

// --- 6. ベンチマーク: go test -bench=. ./reference/07-testing ---
func BenchmarkWordFrequency(b *testing.B) {
	text := "the quick brown fox jumps over the lazy dog the fox runs"
	for b.Loop() {
		WordFrequency(text)
	}
}

// --- 7. Example: ドキュメントになり、Output 行で出力も検証される ---
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("golang"))
	// Output:
	// true
	// false
}
