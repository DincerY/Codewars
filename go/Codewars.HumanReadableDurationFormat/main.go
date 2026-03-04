package main

import "fmt"

func main() {
	FormatDuration(1)
	FormatDuration(3662)
	FormatDuration(62)
	FormatDuration(31622400)
}

func helper(seconds int64) map[string]int64 {
	lib := map[string]int64{
		"years":   0,
		"days":    0,
		"hours":   0,
		"minutes": 0,
		"seconds": 0,
	}
	if seconds >= 60 {
		minutes := seconds / 60
		lib["seconds"] = seconds % 60
		lib["minutes"] = minutes
		if minutes >= 60 {
			hour := minutes / 60
			lib["minutes"] = minutes % 60
			lib["hours"] = hour
			if hour >= 24 {
				days := hour / 24
				lib["hours"] = hour % 24
				lib["days"] = days

				if days >= 365 {
					lib["years"] = days / 365
					lib["days"] = days % 365
				}
			}
		}
	} else {
		if seconds > 0 {
			lib["seconds"] = seconds
		}
	}
	return lib
}

func FormatDuration(seconds int64) string {
	res := ""
	if seconds == 0 {
		return "now"
	}
	lib := helper(seconds)
	count := 0
	for _, v := range lib {
		if v > 0 {
			count++
		}
	}
	str := []string{
		"years",
		"days",
		"hours",
		"minutes",
		"seconds"}
	for _, k := range str {
		v := lib[k]
		if v == 1 {
			if count > 2 {
				res += fmt.Sprintf("%d %s, ", v, k[:len(k)-1])
			}
			if count == 2 {
				res += fmt.Sprintf("%d %s and ", v, k[:len(k)-1])
			}
			if count == 1 {
				res += fmt.Sprintf("%d %s", v, k[:len(k)-1])
			}
			count--
		} else if v > 1 {
			if count > 2 {
				res += fmt.Sprintf("%d %s, ", v, k)
			}
			if count == 2 {
				res += fmt.Sprintf("%d %s and ", v, k)
			}
			if count == 1 {
				res += fmt.Sprintf("%d %s", v, k)
			}
			count--
		}
	}

	return res
}
