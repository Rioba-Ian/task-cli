package helpers

import (
	"fmt"
	"strings"
)

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ExistsInListCmds(slice []string, secondListArg string) bool {
	for _, v := range slice {
		if v == secondListArg {
			return true
		}
	}

	return false
}

func CompareStrings(first, second string) bool {
	f, s := strings.Split(first, " "), strings.Split(second, " ")

	fmt.Printf("f:%s and s:%s will result %t", f, s, f[0] == s[0])

	return f[0] == s[0]
}
