package types

import "fmt"

func EqualFunc[E1, E2 any](s1 []E1, s2 []E2) bool {
    return true
}

func Unexpected() {
    fmt.Println(EqualFunc[int, int]([]int{1}, []int{2}))
}
