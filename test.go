package main
import "fmt"

var (
	a = b
	b = c
	c = f()
)
func f() int {
	return 1
}
func main() {
	fmt.Println(a, b, c)
}