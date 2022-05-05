package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(IsSquare(25)) // Should print 'true'
    fmt.Println(IsSquare(13)) // Should print 'false'
}

func IsSquare(num float64) bool {
    return math.Mod(math.Sqrt(num), 1) == 0
}
