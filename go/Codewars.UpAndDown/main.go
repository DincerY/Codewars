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
	strs := strings.Split(s, " ")

	last := 0

	
	//strs[1], strs[2] = strs[2], strs[1]
	//0,2,4,6,8,10 indexes is lower - (even)
	//1,3,5,7,9,11 indexes is upper - (odd)
	for i := 1; i < len(strs)-1; i++ {	
		if i % 2 == 0 {
			//lower
			if strs[last] < strs[i] {
				strs[last] , strs[i] = strs[i], strs[last] 
			}
			last++
		}else{
			//upper


			last++
		}
	}
	return ""
}