package helper

import (
	"color/src/controller"
	"strings"
)

func GetFlag(args string) (string, string) {

	tester := strings.Split(args, "=")

	if len(tester) == 2 && controller.IsFlag(tester[0]) {
		return tester[0], tester[1]
	}

	return "", args
}
