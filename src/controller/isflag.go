package controller

func IsFlag(expression string) bool {

	flags := []string{"--align", "--output", "--color"}

	for _, flag := range flags {
		if flag == expression {
			return true
		}
	}

	return false
}
