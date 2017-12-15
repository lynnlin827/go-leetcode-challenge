package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
    var count uint = 0
    for m != n {
        m = m >> 1
        n = n >> 1
        count = count + 1
    }
    return m << count
}

func main() {
	fmt.Println(rangeBitwiseAnd(0, 0)) // 0
	fmt.Println(rangeBitwiseAnd(5, 7)) // 4
	fmt.Println(rangeBitwiseAnd(1, 3)) // 0
}
