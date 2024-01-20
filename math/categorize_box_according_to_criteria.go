package math

func categorizeBox(length int, width int, height int, mass int) string {
	bulk := int(1e9)
	dimension := int(1e4)
	quality := 100
	bulky := false
	heavy := false
	if length >= dimension || width >= dimension || height >= dimension {
		bulky = true
	}
	if length*width*height >= bulk {
		bulky = true
	}
	if mass >= quality {
		heavy = true
	}
	if bulky && heavy {
		return "Both"
	}
	if bulky {
		return "Bulky"
	}
	if heavy {
		return "Heavy"
	}
	return "Neither"
}
