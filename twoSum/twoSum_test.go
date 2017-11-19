package twoSum

import (
    "testing"
    "reflect"
)

func TestTwoSum(t *testing.T) {
    result := twoSum([]int{3, 2, 4}, 6);
    if !reflect.DeepEqual(result, []int{1, 2}) {
        t.Errorf("Should be [1, 2]")
    }
}
