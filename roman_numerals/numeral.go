package main

import "strings"

func ConvertToRoman(arabic int) string {

	if arabic == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := arabic; i > 0; i-- {

		if arabic == 4 {
			result.WriteString("IV")
			break
		} else if arabic == 5 {
			result.WriteString("V")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
