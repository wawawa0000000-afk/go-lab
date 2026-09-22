package main
/*
//形式的にはこう書く
import(
	"fmt"
	"strings"
)
*/

import "fmt"
import "strings"
import "unicode/utf8"

func main() {
	var str string
	str = "あ"
	str = str + "い"
	str += "う"
	fmt.Println(str)

	//  str -= "い" <- コンパイルエラーとなる
	//import "strings" の利用
	str = strings.Replace(str, "い", "", 1) // 最初の1つの"い"を削除
	//  				 変数, ”変更前”, "変更後", 個数
	fmt.Println(str)
	str = strings.Replace(str, "あ", "う", 1)
	fmt.Println(str)
	str = strings.Replace(str, "う", "あ", 2)
	fmt.Println(str)
	// 特殊例　個数 = -1
	str = strings.Replace(str, "あ", "い", -1)
	fmt.Println(str) // -1とすることで見つかった変更前をすべて変更後にする

	// 文字列の長さ
	stringlen()
}

func stringlen(){
	var en string = "golang"
	var ja string = "Go言語"

	fmt.Println(en, "len :", len(en)) // ビット単位で出るから 長さ”6”
	fmt.Println(ja, "len :", len(ja)) // バイト単位 長さ”8”
	fmt.Println(ja, "utf8-len :", utf8.RuneCountInString(ja)) // unicode単位で表示 長さ”4”
}

