// Package main - 02: struct とメソッド、ポインタ、埋め込み。
// 実行: go run ./reference/02-structs-methods
package main

import "fmt"

// struct はデータの集まり。クラスではない（継承はない）
type Rectangle struct {
	Width, Height float64
}

// 値レシーバ: コピーに対して動く。読み取り専用の処理向き
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ポインタレシーバ: 元の値を書き換えられる
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// コンストラクタは慣習的に New... 関数で書く（言語機能ではない）
func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{Width: w, Height: h}
}

// --- 埋め込み（composition）: 継承の代わり ---

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " が鳴く"
}

type Dog struct {
	Animal // 匿名フィールド = 埋め込み。Animal のメソッドがそのまま使える
	Breed  string
}

// Dog は Speak を「上書き」できる
func (d Dog) Speak() string {
	return d.Name + " (" + d.Breed + ") がワンと鳴く"
}

func main() {
	fmt.Println("=== struct とメソッド ===")
	r := Rectangle{Width: 3, Height: 4}
	fmt.Printf("%+v  面積=%.1f\n", r, r.Area())

	// ポインタレシーバのメソッドは、値からでも自動で &r される
	r.Scale(2)
	fmt.Printf("Scale(2) 後: %+v  面積=%.1f\n", r, r.Area())

	rp := NewRectangle(10, 2)
	fmt.Printf("NewRectangle: %+v  面積=%.1f\n", *rp, rp.Area())

	fmt.Println("\n=== 値レシーバ vs ポインタレシーバ ===")
	demoReceiver()

	fmt.Println("\n=== 埋め込み ===")
	d := Dog{
		Animal: Animal{Name: "ポチ"},
		Breed:  "柴犬",
	}
	fmt.Println(d.Speak())            // Dog 版
	fmt.Println(d.Animal.Speak())     // 埋め込んだ Animal 版を明示的に呼ぶ
	fmt.Println("名前に直接アクセス:", d.Name) // d.Animal.Name の短縮

	fmt.Println("\n=== struct の比較とコピー ===")
	demoValueSemantics()
}

func demoReceiver() {
	r := Rectangle{Width: 1, Height: 1}

	// 値レシーバのメソッドに渡すとコピーされる
	tryModifyByValue(r)
	fmt.Printf("値渡し後: %+v (変わらない)\n", r)

	// ポインタを渡すと元が変わる
	tryModifyByPointer(&r)
	fmt.Printf("ポインタ渡し後: %+v (変わった)\n", r)
}

func tryModifyByValue(r Rectangle)    { r.Width = 999 }
func tryModifyByPointer(r *Rectangle) { r.Width = 999 }

func demoValueSemantics() {
	a := Rectangle{2, 3}
	b := a // struct は代入で丸ごとコピーされる
	b.Width = 100
	fmt.Printf("a=%+v  b=%+v  (独立している)\n", a, b)

	// 同じ型の struct は == で比較できる（全フィールドが比較可能なら）
	c := Rectangle{2, 3}
	fmt.Println("a == c ?", a == c)
}
