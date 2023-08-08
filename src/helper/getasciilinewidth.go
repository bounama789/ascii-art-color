package helper

func GetAsciiLineWidth(line string, data [][9]string) int {

	result := 0

	for _, char := range line {
		result += len(data[char - 32][0])
	}

	return result
}

func GetAsciiLineWithoutBlankWidth(line string, data [][9]string) int {

	result := 0

	for _, char := range line {
		if char > 32 {
			result += len(data[char - 32][0])
		}
	}

	return result
}
