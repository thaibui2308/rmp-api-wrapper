package h

func SearchInterval(lastname string, initialMap []int) (lower, upper int) {
	if lastname == "" {
		return 0, 0
	}
	asciiVal := GetAsciiCode(lastname)

	for i := 0; i < len(initialMap); i++ {
		if i == len(initialMap)-1 {
			lower = len(initialMap) - 2
			upper = i
		} else {
			if (initialMap[i] <= asciiVal) && (initialMap[i+1] >= asciiVal) {
				lower = i
				upper = i + 1
				break
			}
		}
	}
	return
}
