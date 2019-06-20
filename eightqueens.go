package piscine

import ("fmt")

const numb = 8
var pst = [numb]int{}

func getStatus(qNbr , rPst int) bool{
	for i := 0; i < qNbr; i++  {
		orPst := pst[i]
		if orPst == rPst || orPst == rPst - (qNbr - i) || orPst == rPst + (qNbr - i){
			return false
		}
	}
	return true
}

func resolve(q int){
	if q == numb {
		for i := 0; i < numb; i++ {
			fmt.Print(pst[i] + 1)
		}
		fmt.Print("\n")
	}else {
		for i := 0; i < numb; i++ {
			if getStatus(q, i) {
				pst[q] = i
				resolve(q + 1)
			}
		}
	}
}


func EightQueens() {
	resolve(0)
}
