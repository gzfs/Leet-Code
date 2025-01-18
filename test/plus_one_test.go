package test

import (
	"reflect"
	"testing"

	"github.com/gzfs/Go~Leet/code"
)

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 3}
	incr := code.PlusOne(digits)
	if !reflect.DeepEqual(incr, []int{1, 2, 4}) {
		t.Fatal("Not Equal")
	}

}
