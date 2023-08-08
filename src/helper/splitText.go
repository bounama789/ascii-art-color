package helper

type Paragraph struct {
	Text         string
	NbrOfNewLine int
}

func SplitText(text string) []Paragraph {

	result, line, tabRune := []Paragraph{}, []rune{}, []rune(text)
	lenght := len(tabRune)

	for i := 0; i < lenght; i++ {
		switch tabRune[i] {
		case '\n':
			temp := Paragraph{}

			switch line {
			case nil:
				temp.Text = ""

			default:
				temp.Text = string(line)
				line = nil
			}

			counter := 1
			for i < lenght-1 && tabRune[i+1] == '\n' {
				counter++
				i++
			}
			temp.NbrOfNewLine = counter
			result = append(result, temp)

		default:
			line = append(line, tabRune[i])
		}
	}
	if line != nil {
		temp := Paragraph{}
		temp.Text, temp.NbrOfNewLine = string(line), 0
		result = append(result, temp)
	}

	return result
}
