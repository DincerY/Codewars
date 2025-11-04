package main

import (
	"strings"
	"unicode"
)

func main() {
	ToCamelCase("the-stealth-warrior")
	ToCamelCase("The_Stealth_Warrior")
	ToCamelCase("The_Stealth-Warrior")

}

func ToCamelCase(s string) string {
	var arr []string

	var builder strings.Builder

	for _, chr := range s{
		if unicode.IsLetter(chr) {
			builder.WriteRune(chr)
		}else{
			arr = append(arr, builder.String())
			builder.Reset()
		}
	}
	arr = append(arr, builder.String())

	for i := 1; i < len(arr); i++ {
		r := []rune(arr[i])
		r[0] = unicode.ToTitle(r[0])
		arr[i] = string(r)
	}
	return strings.Join(arr, "")
}