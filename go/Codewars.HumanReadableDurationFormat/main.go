package main

import "fmt"

func main() {
	FormatDuration(31622400)
	FormatDuration(3662)

}

func helper(seconds int64) map[string]int64 {
	lib := map[string]int64{
		"seconds": 0,
		"minutes": 0,
		"hours":   0,
		"days":    0,
		"years":   0,
	}
	if seconds >= 60 {
		minutes := seconds / 60
		lib["seconds"] = seconds % 60
		if minutes >= 60 {
			hour := minutes / 60
			lib["minutes"] = minutes % 60
			if hour >= 24 {
				days := hour / 24
				lib["hours"] = hour % 24

				if days >= 365 {
					lib["years"] = days / 365
					lib["days"] = days % 365
				}
			}
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
	for k, v := range lib {
		if v == 1 {
			res += fmt.Sprintf("%d %s", v, k[:len(k)-1])
			print(res)

		} else if v > 1 {

		}
	}

	return res
}
