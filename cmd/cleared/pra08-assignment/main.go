package main
import "fmt"
func main() {
	/* := で宣言と代入を同時に行う */
	a := 12345
	b := 3*2

	/* 同時に複数の変数に代入可能 */
	c,	d := 3.14, "abcdefg"

	/* 2つの値を返す関数 fn の戻り値を代入する */
	e,	f := fn()

	fmt.Println("int:",a)
	fmt.Println("int:",b)
	fmt.Println("int:",c)
	fmt.Println("string:",d)
	fmt.Println("int:",e)
	fmt.Println("int:",f)
}
func fn()(int,int){
	return 23, 3223
}
