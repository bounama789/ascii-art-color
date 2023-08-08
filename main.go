package main

import (
	"color/src/controller"
	"color/src/helper"
	"color/src/utils"
	"fmt"
	"os"
	"strings"
)

var text, banner, align, output, color, toColor string

func main() {
	text, banner, align, output, color, toColor = SetParams()
	result := utils.PrintAll(text, banner, align, color, toColor)

	if output == "" {
		output = os.Stdout.Name()
	}
	os.WriteFile(output, []byte(result), 0777)
}

// Juste to get and check arguments to set the parameters of PrintAll() function
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>
func SetParams() (string, string, string, string, string, string) {
	argv := os.Args[1:]
	argc := len(argv)

	if argc < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	switch argc {
		case 1:
			flag, _ := helper.GetFlag(argv[0])

			switch flag {
				case "":
					text = strings.ReplaceAll(argv[0], "\\n", "\n")

				default:
					fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
					os.Exit(0)
			}

		case 2:
			flag1, term1 := helper.GetFlag(argv[0])
			flag2, _ := helper.GetFlag(argv[1])

			switch flag1 {
				case "":
					text = strings.ReplaceAll(argv[0], "\\n", "\n")
					banner = argv[1]

				default:
					SetOptions(flag1, term1)

					switch flag2 {
						case "":
							text = strings.ReplaceAll(argv[1], "\\n", "\n")

						default:
							fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
							os.Exit(0)
					}
			}

		case 3:
			flag1, term1 := helper.GetFlag(argv[0])
			flag2, term2 := helper.GetFlag(argv[1])
			flag3, _ := helper.GetFlag(argv[2])

			switch flag1 {
				case "":
					fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
					os.Exit(0)

				default:
					if flag1 == "--color" && flag2 == "" {
						SetOptions(flag1, term1)
						if controller.IsBanner(argv[2]) {
							text = strings.ReplaceAll(argv[1], "\\n", "\n")
							banner = argv[2]
						} else {
							toColor = term2

							switch flag3 {
								case "":
									text = strings.ReplaceAll(argv[2], "\\n", "\n")
		
								default:
									fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
									os.Exit(0)
							}
						}
					} else {
						SetOptions(flag1, term1)
						switch flag2 {
							case "":
								text = strings.ReplaceAll(argv[1], "\\n", "\n")
								banner = argv[2]
			
							default:
								SetOptions(flag2, term2)

								switch flag3 {
									case "":
										text = strings.ReplaceAll(argv[2], "\\n", "\n")
			
									default:
										fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
										os.Exit(0)
								}
						}
					}
			}

			flagTab := []string{flag1, flag2}
			if controller.IsCrash(flagTab) {
				fmt.Print("\n🚨: Options \"--color\" and \"--output\" are incompatible !\n\n")
				os.Exit(0)
			}

		case 4:
			flag1, term1 := helper.GetFlag(argv[0])
			flag2, term2 := helper.GetFlag(argv[1])
			flag3, term3 := helper.GetFlag(argv[2])
			flag4, _ := helper.GetFlag(argv[3])

			switch flag1 {
				case "":
					fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
					os.Exit(0)

				default:
					if flag1 == "--color" && flag2 == "" {
						SetOptions(flag1, term1)
						toColor = term2

						switch flag3 {
							case "":
								text = strings.ReplaceAll(argv[2], "\\n", "\n")
								banner = argv[3]
			
							default:
								SetOptions(flag3, term3)
								switch flag4 {
									case "":
										text = strings.ReplaceAll(argv[3], "\\n", "\n")
			
									default:
										fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
										os.Exit(0)
								}
						}
					} else {
						SetOptions(flag1, term1)
						switch flag2 {
							case "":
								fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
								os.Exit(0)
			
							default:
								if flag2 == "--color" && flag3 == "" {
									SetOptions(flag2, term2)
									if controller.IsBanner(argv[3]) {
										text = strings.ReplaceAll(argv[2], "\\n", "\n")
										banner = argv[3]
									} else {
										toColor = term3
				
										switch flag4 {
											case "":
												text = strings.ReplaceAll(argv[3], "\\n", "\n")
					
											default:
												fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
												os.Exit(0)
										}
									}
								} else {
									SetOptions(flag2, term2)
									switch flag3 {
										case "":
											text = strings.ReplaceAll(argv[2], "\\n", "\n")
											banner = argv[3]
						
										default:
											SetOptions(flag3, term3)
			
											switch flag4 {
												case "":
													text = strings.ReplaceAll(argv[3], "\\n", "\n")
						
												default:
													fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
													os.Exit(0)
											}
									}
								}
						}
					}
			}

			flagTab := []string{flag1, flag2, flag3}
			if controller.IsCrash(flagTab) {
				fmt.Print("\n🚨: Options \"--color\" and \"--output\" are incompatible !\n\n")
				os.Exit(0)
			}

		case 5:
			flag1, term1 := helper.GetFlag(argv[0])
			flag2, term2 := helper.GetFlag(argv[1])
			flag3, term3 := helper.GetFlag(argv[2])
			flag4, term4 := helper.GetFlag(argv[3])
			flag5, _ := helper.GetFlag(argv[4])

			switch flag1 {
				case "":
					fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
					os.Exit(0)

				default:
					if flag1 == "--color" && flag2 == "" {
						SetOptions(flag1, term1)
						toColor = term2

						switch flag3 {
							case "":
								fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
								os.Exit(0)
			
							default:
								if flag3 == "--color" && flag4 == "" {
									SetOptions(flag3, term3)
									if controller.IsBanner(argv[4]) {
										text = strings.ReplaceAll(argv[3], "\\n", "\n")
										banner = argv[4]
									} else {
										toColor = term4
				
										switch flag5 {
											case "":
												text = strings.ReplaceAll(argv[4], "\\n", "\n")
					
											default:
												fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
												os.Exit(0)
										}
									}
								} else {
									SetOptions(flag3, term3)
									switch flag4 {
										case "":
											text = strings.ReplaceAll(argv[3], "\\n", "\n")
											banner = argv[4]
						
										default:
											SetOptions(flag4, term4)
			
											switch flag5 {
												case "":
													text = strings.ReplaceAll(argv[4], "\\n", "\n")
						
												default:
													fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
													os.Exit(0)
											}
									}
								}
						}
					} else {
						SetOptions(flag1, term1)
						switch flag2 {
							case "":
								fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
								os.Exit(0)
			
							default:
								if flag2 == "--color" && flag3 == "" {
									SetOptions(flag2, term2)
									toColor = term3
			
									switch flag4 {
										case "":
											text = strings.ReplaceAll(argv[3], "\\n", "\n")
											banner = argv[4]
						
										default:
											SetOptions(flag4, term4)
											switch flag5 {
												case "":
													text = strings.ReplaceAll(argv[4], "\\n", "\n")
						
												default:
													fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
													os.Exit(0)
											}
									}
								} else {
									SetOptions(flag2, term2)
									switch flag3 {
										case "":
											fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
											os.Exit(0)
						
										default:
											if flag3 == "--color" && flag4 == "" {
												SetOptions(flag3, term3)
												if controller.IsBanner(argv[4]) {
													text = strings.ReplaceAll(argv[3], "\\n", "\n")
													banner = argv[4]
												} else {
													toColor = term4
							
													switch flag5 {
														case "":
															text = strings.ReplaceAll(argv[4], "\\n", "\n")
								
														default:
															fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
															os.Exit(0)
													}
												}
											} else {
												SetOptions(flag3, term3)
												switch flag4 {
													case "":
														text = strings.ReplaceAll(argv[3], "\\n", "\n")
														banner = argv[4]
									
													default:
														SetOptions(flag4, term4)
						
														switch flag5 {
															case "":
																text = strings.ReplaceAll(argv[4], "\\n", "\n")
									
															default:
																fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
																os.Exit(0)
														}
												}
											}
									}
								}
						}
					}
			}

			flagTab := []string{flag1, flag2, flag3, flag4}
			if controller.IsCrash(flagTab) {
				fmt.Print("\n🚨: Options \"--color\" and \"--output\" are incompatible !\n\n")
				os.Exit(0)
			}

		default:
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
	}

	return text, banner, align, output, color, toColor
}

// Juste to set value(s) for the option(s)
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>
func SetOptions(flag, term string) {
	if flag == "--align" && align != "" || flag == "--output" && output != "" || flag == "--color" && color != "" {
		fmt.Print("\n🚨: An option entered more than 1 time !\n\n")
		os.Exit(0)
	} else {
		if flag == "--align" {
			align = term
		} else if flag == "--output" {
			output = term
		} else {
			color = term
		}
	}
}
