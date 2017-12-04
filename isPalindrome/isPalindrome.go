package main

import "strings"
import "fmt"

func isPalindrome(s string) bool {
    r := []byte(strings.ToLower(s))
    len := len(r)
    i := 0
    j := len - 1
    for i < len {
        if !isAlphaNumeric(r[i]) {
            i = i + 1
            continue
        }
        if !isAlphaNumeric(r[j]) {
            j = j - 1
            continue
        }
        if r[i] != r[j] {
            return false
        }
        i = i + 1
        j = j - 1
    }
    return true
}

func isAlphaNumeric(b byte) bool {
    if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
        return true
    }
    return false
}

func main() {
    fmt.Println(isPalindrome("raceacar")) // false
    fmt.Println(isPalindrome("AmanaplanacanalPanama")) // true
    fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
    fmt.Println(isPalindrome("0P")) // false
}
