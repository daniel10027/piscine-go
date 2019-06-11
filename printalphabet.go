package main

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