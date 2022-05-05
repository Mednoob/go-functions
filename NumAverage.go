package main

import (
    "fmt"
)

func main() {
    arr := []float64{4, 5, 3}
    fmt.Println(Average(arr)) // Should print '4'
    arr = append(arr, 3.1)
    fmt.Println(Average(arr)) // Should print '3.775'
}

func Average(arr []float64) float64 {
    var total float64 = 0

    for _, v := range arr {
        total += v
    }

    return total / float64(len(arr))
}
