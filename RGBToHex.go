package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

func main() {
    fmt.Println(RGBToHex(255, 255, 255)) // Should print 'FFFFFF'
    fmt.Println(RGBToHex(148, 0, 211)) // Should print '9400D3'
}

func RGBToHex(r, g, b float64) string {
    return Hex(r) + Hex(g) + Hex(b)
}

func Hex(num float64) string {
    first := math.Max(num, 0)
    first = math.Min(num/16, 15.9375)
    second := math.Mod(first, 1) * 16

    return strings.ToUpper(strconv.FormatInt(int64(math.Floor(first)), 16) + strconv.FormatInt(int64(second), 16))
}
