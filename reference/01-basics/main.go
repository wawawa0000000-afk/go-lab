// Package main - 01: Go の基本文法をひと通り触る。
// 実行: go run ./reference/01-basics
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== 1. 変数と型 ===")
	variablesAndTypes()

	fmt.Println("\n=== 2. 制御構文 ===")
	controlFlow()

	fmt.Println("\n=== 3. 関数（複数戻り値・可変長） ===")
	functions()

	fmt.Println("\n=== 4. スライス ===")
	slices()

	fmt.Println("\n=== 5. マップ ===")
	maps()

	fmt.Println("\n=== 6. エラー処理 ===")
	errorHandling()

	fmt.Println("\n=== 7. defer ===")
	deferDemo()
}

func variablesAndTypes() {
	var a int = 10        // 明示的な型
	b := 3.14             // 型推論（:= は関数内でのみ使える）
	const greeting = "hi" // 定数
	var flag bool         // ゼロ値 = false

	// Go は暗黙の型変換をしない。明示的に変換する
	sum := float64(a) + b

	fmt.Printf("a=%d b=%.2f greeting=%q flag=%t sum=%.2f\n", a, b, greeting, flag, sum)
	fmt.Printf("型: a=%T b=%T\n", a, b)
}

func controlFlow() {
	// if には初期化文を書ける。スコープは if 内に限定される
	if n := 7; n%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数:", n)
	}

	// for は Go 唯一のループ構文。while 相当も for で書く
	total := 0
	for i := 1; i <= 5; i++ {
		total += i
	}
	fmt.Println("1..5 の合計:", total)

	// range でスライスを回す
	for idx, ch := range "Go言語" {
		fmt.Printf("  index=%d rune=%c\n", idx, ch)
	}

	// switch は break 不要。fallthrough は明示的に書く
	switch day := 3; day {
	case 1, 2, 3, 4, 5:
		fmt.Println("平日")
	case 6, 7:
		fmt.Println("週末")
	}
}

// 複数の値を返せる。慣習として最後の戻り値が error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ゼロ除算")
	}
	return a / b, nil
}

// 可変長引数
func sum(nums ...int) int {
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}

func functions() {
	q, err := divide(10, 3)
	fmt.Println("10 / 3 =", q, "err:", err)

	_, err = divide(1, 0)
	fmt.Println("1 / 0 err:", err)

	fmt.Println("sum(1,2,3,4):", sum(1, 2, 3, 4))

	// 無名関数・クロージャ
	counter := makeCounter()
	fmt.Println(counter(), counter(), counter()) // 1 2 3
}

func makeCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func slices() {
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	// スライスの一部を取り出す（元の配列を共有する点に注意）
	sub := s[1:3]
	fmt.Println("s[1:3]:", sub)

	// 2次元
	grid := make([][]int, 2)
	for i := range grid {
		grid[i] = make([]int, 3)
	}
	grid[0][1] = 9
	fmt.Println("grid:", grid)
}

func maps() {
	m := map[string]int{"apple": 3, "banana": 5}
	m["cherry"] = 7

	// キーの存在チェック（カンマ ok イディオム）
	if v, ok := m["apple"]; ok {
		fmt.Println("apple:", v)
	}
	if _, ok := m["grape"]; !ok {
		fmt.Println("grape は無い")
	}

	delete(m, "banana")
	fmt.Println("map:", m)
}

// 独自エラー型
type NotFoundError struct {
	Key string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("キー %q が見つからない", e.Key)
}

func lookup(key string) (string, error) {
	data := map[string]string{"jp": "日本", "us": "アメリカ"}
	v, ok := data[key]
	if !ok {
		return "", &NotFoundError{Key: key}
	}
	return v, nil
}

func errorHandling() {
	v, err := lookup("jp")
	fmt.Println("jp:", v, err)

	_, err = lookup("xx")
	fmt.Println("xx err:", err)

	// errors.As で型を取り出す
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		fmt.Println("  → NotFoundError だった。Key =", nfe.Key)
	}

	// エラーのラップ
	wrapped := fmt.Errorf("処理失敗: %w", err)
	fmt.Println("wrapped:", wrapped)
	fmt.Println("errors.As は wrap 越しでも効く:", errors.As(wrapped, &nfe))
}

func deferDemo() {
	// defer は関数の終了時に LIFO で実行される。後片付けに使う
	defer fmt.Println("  defer 3 (最後に実行)")
	defer fmt.Println("  defer 2")
	defer fmt.Println("  defer 1 (最初に登録 → 最後から3番目)")
	fmt.Println("  本体:", strings.Repeat("-", 10))
}
