package test

import (
	"testing"

	"github.com/gzfs/Go~Leet/code"
)

func TestSqrt(t *testing.T) {
    if code.Sqrt(8) != 2 {
        t.Fatalf("Error in SQRT");
    } 
}
