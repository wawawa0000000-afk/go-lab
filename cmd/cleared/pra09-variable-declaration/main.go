package main

import "fmt"

func main() {
	// 変数xの宣言
	var x int

	// 同じ型の変数を複数宣言する際、変数名をカンマを使って列挙する
	var a, b, c int

	// 丸括弧を使ったグルーピング
	var (
		i      int
		s1, s2 string
	)

	fmt.Println("x:", x)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("i:", i)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	syouryaku()
}

func syouryaku() {
	var x = 1

	var a, b, c = 1, 2, 3

	var (
		i      = 123
		s1, s2 = "aaa", "bbb"
	)

	fmt.Println("x:", x)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("i:", i)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	tyokusetu()
}

func tyokusetu() {
	x := 123

	a, b, c := 1, 2, 3

	i := 1
	s1, s2 := "aaa", "bbb"

	fmt.Println("x:", x)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("i:", i)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

}
