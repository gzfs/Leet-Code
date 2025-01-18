package test

import (
	"testing"

	"github.com/gzfs/Go~Leet/code"
)

func TestBinarySearch(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6};
    target := 4;

    index := code.BinarySearch(nums, target);

    non_existent_target := 10;
    negative_index := code.BinarySearch(nums, non_existent_target);

    if negative_index != -1 {
        t.Fatal("Positive Index returned for a Non Existent Element");
    }

    if index != 3 {
        t.Fatal("Incorrect Index returned for an Existing Element");
    }
}
