package style

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SetColor(color string) string {

	switch color[0] {
		case '#':
			hexFormat := regexp.MustCompile(`^[0-9a-fA-F]+$`)
			color = strings.TrimPrefix(color, "#")

			switch hexFormat.MatchString(color) {
				case true:
					return RGBtoANSI(HEXtoRGB(color))

				default:
					fmt.Fprintf(os.Stderr, "\n🚨: `#%s` is not an hexadecimal number\n\n", color)
					os.Exit(0)
			}

		default:
			RGBformat := regexp.MustCompile(`^rgb\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)$`)
			match := RGBformat.FindStringSubmatch(color)

			if len(match) == 4 {
				// Convert the captured strings to integers
				red, _ := strconv.Atoi(match[1])
				green, _ := strconv.Atoi(match[2])
				blue, _ := strconv.Atoi(match[3])

				if red < 256 && green < 256 && blue < 256 {
					return RGBtoANSI(red, green, blue)
				}
			}

			colorRgb := COLOR[color]
	
			if colorRgb != nil {
				return RGBtoANSI(colorRgb[0], colorRgb[1], colorRgb[2])
			}
	}

	return ""
}

func RGBtoANSI(red, green, blue int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", red, green, blue)
}

func HEXtoRGB(color string) (int, int, int) {

	red, _ := strconv.ParseInt(color[0:2], 16, 0)
	green, _ := strconv.ParseInt(color[2:4], 16, 0)
	blue, _ := strconv.ParseInt(color[4:6], 16, 0)

	return int(red), int(green), int(blue)
}
