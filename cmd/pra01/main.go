package main
import "fmt"

func start() int64 {
	i := 12345
	var i64 int64 = int64(i)

	return i64
}
func next() string {
	word := "sandwich"

	return word
}
func main(){
	id := start() 
	name := next()
	fmt.Println("id:",id, "name:",name)
}
