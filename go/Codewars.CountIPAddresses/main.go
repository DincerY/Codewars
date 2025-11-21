package main

import (
	"encoding/binary"
	"math"
	"net"
	"strconv"
	"strings"
)

func main() {
	IpsBetween2("10.0.0.0", "10.0.0.50")
	IpsBetween("10.0.0.0", "10.0.1.0")
	IpsBetween("20.0.0.10", "20.0.1.0")
}

func IpsBetween(start, end string) uint32 {
	var total int32
	startArr := strings.Split(start, ".")
	endArr := strings.Split(end, ".")

	for i := 3; i >= 0; i-- {
		val1, _ := strconv.Atoi(endArr[i])
		val2, _ := strconv.Atoi(startArr[i])

		total += int32((val1 - val2) * int(math.Pow(float64(256), float64(3-i))))
	}

	return uint32(total)
}

func IpsBetween2(start, end string) uint32 {
	startVal := binary.BigEndian.Uint32(net.ParseIP(start)[12:16])
	endVal := binary.BigEndian.Uint32(net.ParseIP(end)[12:16])

	return startVal - endVal
}
