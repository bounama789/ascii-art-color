package utils

import "color/src/helper"

func PrintAll(text, banner, align, color, toColor string) string {

	data := helper.SplitText(text)
	lenght, output := len(data), ""

	if banner == "" {
		banner = "standard"
	}
	if align == "" {
		align = "left"
	}

	for i, ln := range data {
		if ln.Text != "" {
			output += ToAsciiLine(ln.Text, banner, align, color, toColor)
		}

		k := 0
		if i < lenght-1 && ln.Text != "" {
			k = 1
		}
		for ; k < ln.NbrOfNewLine; k++ {
			output += "\n"
		}
	}

	return output
}
