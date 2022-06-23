package main

import "fmt"

func main() {
    arr := []string{}
    fmt.Println(HumanJoin(arr)) // Should print ''

    arr = append(arr, "a")
    fmt.Println(HumanJoin(arr)) // Should print 'a'

    arr = append(arr, "b")
    fmt.Println(HumanJoin(arr)) // Should print 'a and b'

    arr = append(arr, "c")
    fmt.Println(HumanJoin(arr)) // Should print 'a, b, and c'

    arr = append(arr, "d")
    fmt.Println(HumanJoin(arr)) // Should print 'a, b, c, and d'
}

func HumanJoin(arr []string) string {
    res := ""
    arrLen := len(arr)

    for i, v := range arr {
        res += v
        if i != arrLen-1 {
            if arrLen > 2 {
                res += ", "
            } else {
                res += " "
            }
        }
        if i == arrLen-2 {
            res += "and "
        }
    }

    return res
}
