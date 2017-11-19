package reverseInteger

import "testing"

func TestReverse(t *testing.T) {
    if result := reverse(123); result != 321 {
        t.Errorf("Should be 321")
    }
    if result := reverse(-123); result != -321 {
        t.Errorf("Should be -321")
    }
    if result := reverse(120); result != 21 {
        t.Errorf("Should be 21")
    }
    if result := reverse(201); result != 102 {
        t.Errorf("Should be 102")
    }
    if result := reverse(1534236469); result != 0 {
        t.Errorf("Should be 0")
    }
    if result := reverse(-2147483648); result != 0 {
        t.Errorf("Should be 0")
    }
}
