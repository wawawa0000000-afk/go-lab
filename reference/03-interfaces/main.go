// Package main - 03: インターフェース。Go の多態はこれで実現する。
// 実行: go run ./reference/03-interfaces
package main

import (
	"fmt"
	"math"
	"strings"
)

// インターフェース = メソッドの集合。
// 「implements」を書く必要はない。メソッドを満たせば自動的にその型になる（構造的型付け）。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ R float64 }

func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

type Square struct{ Side float64 }

func (s Square) Area() float64      { return s.Side * s.Side }
func (s Square) Perimeter() float64 { return 4 * s.Side }

// Shape を受け取る関数は、Circle でも Square でも動く
func describe(s Shape) {
	fmt.Printf("  %T → 面積=%.2f 周=%.2f\n", s, s.Area(), s.Perimeter())
}

func main() {
	fmt.Println("=== 多態 ===")
	shapes := []Shape{
		Circle{R: 2},
		Square{Side: 3},
	}
	for _, s := range shapes {
		describe(s)
	}

	fmt.Println("\n=== 型アサーションと type switch ===")
	for _, s := range shapes {
		switch v := s.(type) {
		case Circle:
			fmt.Printf("  円: 半径 %.1f\n", v.R)
		case Square:
			fmt.Printf("  正方形: 一辺 %.1f\n", v.Side)
		}
	}

	// 単発の型アサーション（カンマ ok）
	var s Shape = Circle{R: 1}
	if c, ok := s.(Circle); ok {
		fmt.Printf("  これは Circle。R=%.1f\n", c.R)
	}

	fmt.Println("\n=== 空インターフェース any ===")
	printAny(42)
	printAny("hello")
	printAny([]int{1, 2, 3})

	fmt.Println("\n=== 標準インターフェースを実装する ===")
	// fmt.Stringer を満たすと Println で使われる
	t := Temperature(25.5)
	fmt.Println("  温度:", t)

	// io.Writer を満たす自作型
	var w strings.Builder
	fmt.Fprintf(&w, "Fprintf は io.Writer になら何にでも書ける")
	fmt.Println("  builder:", w.String())

	fmt.Println("\n=== インターフェースのゼロ値は nil ===")
	var empty Shape
	fmt.Println("  empty == nil ?", empty == nil)
}

func printAny(v any) { // any は interface{} の別名（Go 1.18+）
	fmt.Printf("  値=%v 型=%T\n", v, v)
}

// fmt.Stringer インターフェース: String() string を持てばよい
type Temperature float64

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", float64(t))
}
