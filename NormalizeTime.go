package main

import (
    "fmt"
    "math"
    "strconv"
)

func main() {
    num := 3660
    fmt.Println(NormalizeTime(num)) // Should print '01:01:00'
    num += 81
    fmt.Println(NormalizeTime(num)) // Should print '01:02:21'
}

func NormalizeTime(second int) string {
    hours := int(math.Floor(float64(second / 3600)))
    minutes := int(math.Floor(float64((second % 3600) / 60)))
    seconds := second % 3600 % 60

    return FixedNum(hours) + ":" + FixedNum(minutes) + ":" + FixedNum(seconds)
}

func FixedNum(num int) string {
    str := strconv.Itoa(num)
    if len(str) < 2 {
        str = "0" + str
    }

    return str
}
