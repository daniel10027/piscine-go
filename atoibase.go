package piscine

func AtoiBase(s string,base string) int{//convertir un int dans une base et retourne le string de celui-ci
	if len(base)<2 || !uniquealphaandnosigne(base){
		return 0
	}else{
		result:=0
		deb:=0
		mul:=1
		if s[0]==43 || s[0]==45{
			deb=1
			if s[0]==45{
				mul=-1
			}
		}
		for i:=deb;i<len(s);i++{
			result+= Index(base,string(s[i]))*RecursivePower(len(base),len(s)-i-1)
		}
		return result*mul
	}
}

