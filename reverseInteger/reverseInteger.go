package reverseInteger

import "math"

func reverse(x int) int {
    result := 0
    for x != 0 {
        remainder := x % 10
        result = result * 10 + remainder
        x = x / 10
    }
    if result > math.MaxInt32 || result < -math.MaxInt32 {
        result = 0
    }
    return result
}
