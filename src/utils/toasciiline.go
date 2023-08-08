package utils

import (
	"color/src/controller"
	"color/src/helper"
	"color/src/style"
	"fmt"
	"os"
)

func ToAsciiLine(text, banner, align, color, toColor string) string {

	tab, length := helper.RangeIndex(text)
	data, result := Init(banner), ""
	jTab := helper.SplitLine(text)

	tWidth := helper.GetTerminalWidth()
	lWidth := helper.GetAsciiLineWidth(text, data)
	blank := tWidth - lWidth

	if blank <= 0 {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	if align == "justify" && len(jTab) == 1 {
		align = "left"
	}

	switch align {
		case "center":
			half := blank / 2
			for i := 0; i < 9; i++ {
				for k := 0; k < half; k++ {
					result += " "
				}
				for j := 0; j < length; j++ {
					switch color {
						case "":
							result += data[tab[j]][i]

						default:
							switch toColor {
								case "":
									result += style.SetColor(color) + data[tab[j]][i] + "\033[0m"

								default:
									switch controller.IsToColor(tab[j], toColor) {
										case true:
											result += style.SetColor(color) + data[tab[j]][i] + "\033[0m"
				
										default:
											result += data[tab[j]][i]
									}
							}
					}
				}
				result += "\n"
			}

		case "left":
			for i := 0; i < 9; i++ {
				for j := 0; j < length; j++ {
					switch color {
						case "":
							result += data[tab[j]][i]

						default:
							switch toColor {
								case "":
									result += style.SetColor(color) + data[tab[j]][i] + "\033[0m"

								default:
									switch controller.IsToColor(tab[j], toColor) {
										case true:
											result += style.SetColor(color) + data[tab[j]][i] + "\033[0m"
				
										default:
											result += data[tab[j]][i]
									}
							}
					}
				}
				result += "\n"
			}

		case "right":
			for i := 0; i < 9; i++ {
				for k := 0; k < blank; k++ {
					result += " "
				}
				for j := 0; j < length; j++ {
					switch color {
						case "":
							result += data[tab[j]][i]

						default:
							switch toColor {
								case "":
									result += style.SetColor(color) + data[tab[j]][i] + "\033[0m"

								default:
									switch controller.IsToColor(tab[j], toColor) {
										case true:
											result += style.SetColor(color) + data[tab[j]][i] + "\033[0m"
				
										default:
											result += data[tab[j]][i]
									}
							}
					}
				}
				result += "\n"
			}

		case "justify":
			jLength := len(jTab)
			first := jTab[0]
			jTab = jTab[1:]
			jWidth := helper.GetAsciiLineWithoutBlankWidth(text, data)
			distance := (tWidth - jWidth) / (jLength - 1)
			rest := blank % (jLength - 1)

			for i := 0; i < 9; i++ {
				for _, char := range first {
					switch color {
						case "":
							result += data[int(char - 32)][i]

						default:
							switch toColor {
								case "":
									result += style.SetColor(color) + data[int(char - 32)][i] + "\033[0m"
	
								default:
									switch controller.IsToColor(int(char - 32), toColor) {
										case true:
											result += style.SetColor(color) + data[int(char - 32)][i] + "\033[0m"
				
										default:
											result += data[int(char - 32)][i]
									}
							}
					}
				}

				for j := 0; j < jLength-1; j++ {
					for k := 0; k < distance; k++ {
						result += " "
					}

					if j < rest {
						result += " "
					}

					for _, char := range jTab[j] {
						switch color {
							case "":
								result += data[int(char - 32)][i]
	
							default:
								switch toColor {
									case "":
										result += style.SetColor(color) + data[int(char - 32)][i] + "\033[0m"
		
									default:
										switch controller.IsToColor(int(char - 32), toColor) {
											case true:
												result += style.SetColor(color) + data[int(char - 32)][i] + "\033[0m"
					
											default:
												result += data[int(char - 32)][i]
										}
								}
						}
					}
				}

				result += "\n"
			}

		default:
			fmt.Print("\n🚨 \"--align\" OPTION : ["+ align +"] value not supported\n\n")
			os.Exit(0)
	}

	return result
}
