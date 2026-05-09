package calc

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

func main() {
	fmt.Println(Add(1, 2))
}
