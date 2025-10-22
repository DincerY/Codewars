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
	if len(listArt) == 0 || len(listCat) == 0{
		return ""
	}
	resMap := make(map[string]int)

	for _, art := range listArt {
		splitted := strings.Split(art, " ")
		firstChar := string(splitted[0][0])
		i, _ := strconv.Atoi(splitted[1])
		resMap[firstChar]+=i
		
	}

	result := make([]string, len(listCat))
	for i, code := range listCat{
		result[i] = fmt.Sprintf("%s : %v", code, resMap[code])
	}
	return strings.Join(result, " - ")
}
