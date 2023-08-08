package utils

import (
	"bufio"
	"color/src/controller"
	"fmt"
	"os"
)

func Init(banner string) [][9]string {

	if !controller.IsBanner(banner) {
		fmt.Print("\n🚨: banner [" + banner + "] not supported\n\n")
		os.Exit(0)
	}

	file, _ := os.Open("assets/" + banner + ".txt")
	reader := bufio.NewScanner(file)
	result, letter, counter, line := [][9]string{}, [9]string{}, 0, ""

	for reader.Scan() {
		line = reader.Text()
		switch line {
		case "":
			if counter > 0 {
				result = append(result, letter)
				letter = [9]string{}
				counter = 0
			}

		default:
			letter[counter] = line
			counter++
		}
	}
	result = append(result, letter)

	file.Close()
	return result
}
