package controller

func IsCrash(flagTab []string) bool {

	counter := 0

	for _, flag := range flagTab {
		if flag == "--output" || flag == "--color" {
			counter++
		}
	}

	return counter > 1
}
