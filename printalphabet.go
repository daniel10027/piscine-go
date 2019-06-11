package main
import "github.com/01-edu/O1"
func printAlphabet() {
	i := byte('a')
	for i <= 'z' {
		O1.Putchar(i)
		i++
	}
}

func main() {
	printAlphabet()
}
