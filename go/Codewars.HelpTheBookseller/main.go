package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	println(StockList(b, c))
}

func StockList(listArt []string, listCat []string) string {
	resMap := make(map[byte]int)
	for _, cat := range listCat {
		resMap[cat[0]] = 0
	}

	for _, art := range listArt {
		if _, ok := resMap[art[0]]; ok {
			value := strings.Split(art, " ")[1]
			i, _ := strconv.Atoi(value)
			resMap[art[0]]+=i
		}
	}
	var builder strings.Builder
	for key, val := range resMap{
		builder.WriteString(fmt.Sprintf("(%c : %d)", key, val))
	}
	return builder.String()
}
