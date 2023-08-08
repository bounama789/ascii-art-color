package helper

import (
	"fmt"
	"os"
)

func RangeIndex(text string) ([]int, int) {

	result, length := []int{}, 0

	for _, r := range text {
		if r < 32 || r > 126 {
			fmt.Print("\n🚨: no ASCII format available for character '"+ string(r) +"'\n\n")
			os.Exit(0)
		}
		result = append(result, int(r - 32))
		length++
	}

	return result, length
}
