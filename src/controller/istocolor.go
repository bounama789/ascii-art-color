package controller

func IsToColor(char int, toColor string) bool {

	chr := rune(char + 32)

	for _, rne := range toColor {
		if rne == chr {
			return true
		}
	}

	return false
}
