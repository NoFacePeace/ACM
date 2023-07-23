package str

func squareIsWhite(coordinates string) bool {
	if len(coordinates) != 2 {
		return false
	}
	num := int(coordinates[0]-'a') + int(coordinates[1]-'0')
	if num%2 == 0 {
		return true
	}
	return false
}
