package main

import (
	"strings"
)

func main() {
	Arrange("who hit retaining The That a we taken")
	Arrange("on I came up were so grandmothers")
	Arrange("way the my wall them him")
}

func Arrange(s string) string {
	strs := strings.Fields(s)
	
	//strs[1], strs[2] = strs[2], strs[1]
	//0,2,4,6,8,10 indexes is lower - (even)
	//1,3,5,7,9,11 indexes is upper - (odd)

	for i := 0; i < len(strs)-1; i++ {	
		if i % 2 == 0 {
			if len(strs[i]) > len(strs[i+1]) {
				strs[i] , strs[i+1] = strs[i+1], strs[i] 
			}
		}else{
			if len(strs[i]) < len(strs[i+1]) {
				strs[i] , strs[i+1] = strs[i+1], strs[i] 
			}
		}
	}
	for i,_ := range strs{
		if i % 2 == 1 {
			strs[i] = strings.ToUpper(strs[i])
			
		}else{
			strs[i] = strings.ToLower(strs[i])
		}
	}
	return strings.Join(strs, " ")
}