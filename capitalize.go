package piscine

func Capitalize(s string) string {
	sRune := []rune(s)
	distance := 'a' - 'A'
	if (sRune[0] >= 97 && sRune[0] <= 122 ) {
		sRune[0] = sRune[0] - distance
	}

	for i := 1; i < len(sRune); i++ {
		if (sRune[i] >= 97 && sRune[i] <= 122) && !(97 <= sRune[i-1] && sRune[i-1] <= 122) && !(65 <= sRune[i-1] && sRune[i-1] <= 90) && !(48 <= sRune[i-1] && sRune[i-1] <= 57){
			sRune[i] = sRune[i] - distance
		}else if ((97 <= sRune[i-1] && sRune[i-1] <= 122) || (65 <= sRune[i-1] && sRune[i-1] <= 90) || (48 <= sRune[i-1] && sRune[i-1] <= 57)) && (sRune[i] >= 65 && sRune[i] <= 90) {
			sRune[i] = sRune[i] + distance
		}
	}
	
	return string(sRune)
}
