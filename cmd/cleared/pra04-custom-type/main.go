package main
import "fmt"

func main() {
	type myInteger int
	var i myInteger = 123
	
	i += 1
	fmt.Println(i)

	type myStruct struct{
		a int 
		b int
	}

	str1 := myStruct{a: 2,b: 6}
	str2 := myStruct{5, 11}
	var str3 myStruct
		str3.a = 12
		//str3.b = 21
	str4 := &myStruct{a: 16}
	//ゼロ値: 各フィールドがその型のゼロ値
	// （string は "", int は 0, bool は false）
	// nil にはならない

	fmt.Println(str1.a, ":", str1.b)
	fmt.Println(str2.a, ":", str2.b)
	fmt.Println(str3.a, ":", str3.b)
	fmt.Println(str4.a, ":", str4.b)
	changer()
}

func changer(){
	var i int  = 1234
	var u uint32 = uint32(i)
	var f float32 = float32(u)
	var s string = string(rune(i))
	var b []byte = []byte("abc")

	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}
