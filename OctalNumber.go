package main

import (
    "fmt"
    "math"
    "strconv"
)

func main() {
    // Octal number of 321 = 3 * 8^2 + 2 * 8^1 + 1 * 8^0 = 209
    fmt.Println(OctalNumber(321))
}

func OctalNumber(num int) float64 {
    var octal float64
    str := strconv.Itoa(num)
    for i, v := range str {
        conv, er := strconv.Atoi(string(v))
        if er != nil {
            fmt.Println("Error occured while converting ASCII to Int")
            return 0
        }
        octal += float64(conv) * math.Pow(8, float64(len(str)-1-i))
    }

    return octal
}
