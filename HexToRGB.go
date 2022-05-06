package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(HexToRGB("FFFFFF")) // Should print '[255 255 255]'
    fmt.Println(HexToRGB("9400D3")) // Should print '[148 0 211]'
}

func HexToRGB(str string) [3]int {
    if len(str) != 6 {
        return [3]int{0, 0, 0}
    }

    return [3]int{RGBValue(str[0:2]), RGBValue(str[2:4]), RGBValue(str[4:6])}
}

func RGBValue(str string) int {
    if len(str) != 2 {
        return 0
    }

    a, er := strconv.ParseInt(string(str[0]), 16, 0)
    b, err := strconv.ParseInt(string(str[1]), 16, 0)

    if er != nil || err != nil {
        return 0
    }

    return int((float32(a) + (float32(b) / 16)) * 16)
}
