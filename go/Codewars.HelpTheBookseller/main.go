package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	
	b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	c = []string{"A", "B"}
	println(StockList(b, c))
}

func StockList(listArt []string, listCat []string) string {
	resMap := make(map[string]int)
	for _, cat := range listCat {
		resMap[cat] = 0
	}

	for _, art := range listArt {
		if _, ok := resMap[art]; ok {
			value := strings.Split(art, " ")[1]
			i, _ := strconv.Atoi(value)
			resMap[art]+=i
		}
	}
	var builder strings.Builder

	for _,cat := range listCat{
		if cat != listCat[0] {
			builder.WriteString(" - ")
		}
		builder.WriteString(fmt.Sprintf("(%c : %d)", cat[0],resMap[cat]))
	}
	return builder.String()

	//Different Approach
	/*result := make([]string, len(listCat))
	for i, code := range listCat{
		result[i] = fmt.Sprintf("%s : %v", code, resMap[code])
	}
	return strings.Join(result, " - ")*/
}
