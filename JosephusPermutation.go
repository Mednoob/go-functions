package main

import "fmt"

func main() {
    arr := make([]interface{}, 0)
    arr = append(arr, 1, 2, 3, 4, 5, 6, 7)
    fmt.Println(JosephusPermutation(arr, 3)) // Should print '[3 6 2 7 5 1 4]'
    arr = make([]interface{}, 0)
    arr = append(arr, "a", "b", "c", "d", "e", "f", "g", "h")
    fmt.Println(JosephusPermutation(arr, 5)) // Should print '[e b h g a d f c]'
}

func JosephusPermutation(items []interface{}, k int) []interface{} {
    clone := make([]interface{}, 0)
    clone = append(clone, items...)
    res := make([]interface{}, 0)
    last := 0
    for len(clone) > 0 {
        index := (last - 1 + k) % len(clone)
        res = append(res, clone[index])
        clone = RemoveIndex(clone, index)
        last = index
    }

    return res
}

func RemoveIndex(s []interface{}, index int) []interface{} {
    ret := make([]interface{}, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}
