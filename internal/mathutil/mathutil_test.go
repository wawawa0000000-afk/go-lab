package mathutil

import (
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {
	// テーブル駆動テスト（Go の定番スタイル）
	cases := []struct {
		name string
		a, b int64
		want int64
	}{
		{"共通因数あり", 12, 18, 6},
		{"互いに素", 17, 5, 1},
		{"片方が0", 0, 9, 9},
		{"負の値", -12, 18, 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := GCD(tc.a, tc.b); got != tc.want {
				t.Errorf("GCD(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestFibSequence(t *testing.T) {
	if got := FibSequence(0); len(got) != 0 {
		t.Errorf("FibSequence(0) = %v, want empty", got)
	}

	want := []int64{0, 1, 1, 2, 3, 5, 8, 13}
	if got := FibSequence(8); !reflect.DeepEqual(got, want) {
		t.Errorf("FibSequence(8) = %v, want %v", got, want)
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(1071, 462)
	}
}
