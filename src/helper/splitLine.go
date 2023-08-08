package helper

func SplitLine(line string) []string {

	result := []string{}
	var word []rune = nil

	for _, char := range line {
		switch char {
			case ' ':
				if word != nil {
					result = append(result, string(word))
					word = nil
				}

			default:
				word = append(word, char)
		}
	}
	if word != nil {
		result = append(result, string(word))
	}

	return result
}
