package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := HexadecimalNumber("fa0")
    fmt.Println(num) // Should print '4000'
    fmt.Println(num == 0xfa0) // Should print 'true'
}

func HexadecimalNumber(str string) int {
    res := 0

    for _, v := range str {
        num, err := strconv.ParseInt(string(v), 16, 0)

        if err != nil {
            return 0
        }

        res = (res * 16) + int(num)
    }

    return res
}
