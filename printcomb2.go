package piscine

import "fmt"

func PrintComb2() {
	for a:=0;a<=97;a++{
		for b:=0;b<=99;b++{
			if a<b{
				fmt.Printf("%02d %02d, ",a,b)
			}
		}
	}
	fmt.Printf("98 99\n")
				
}
