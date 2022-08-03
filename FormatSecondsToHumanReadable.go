package main

import (
    "fmt"
    "strconv"
)

var times = map[string]int{
    "year":   365 * 24 * 60 * 60,
    "day":    24 * 60 * 60,
    "hour":   60 * 60,
    "minute": 60,
    "second": 1,
}

func main() {
	fmt.Println(FormatSecondsToHumanReadable(120))  // Should print '2 minutes'

    fmt.Println(FormatSecondsToHumanReadable(7322)) // Should print '2 hours, 2 minutes and 2 seconds'
    // Note: Sometimes, the result of the above can be "122 minutes and 2 seconds" instead.
    // This is probably because the key loop is not in the same order as in the map declaration
}

func FormatSecondsToHumanReadable(seconds int) string {
    if seconds <= 0 {
        return "now"
    }

    keys := []string{}
    nums := []int{}

    for k := range times {
        if seconds >= times[k] {
            keys = append(keys, k)
            nums = append(nums, seconds/times[k])
            seconds %= times[k]
        }
    }

    res := ""

    for i, key := range keys {
        num := nums[i]

        res += strconv.Itoa(num) + " " + key

        if num > 1 {
            res += "s"
        }

        if i == len(keys)-2 {
            res += " and "
        } else if i < len(keys)-2 {
            res += ", "
        }
    }

    return res
}
