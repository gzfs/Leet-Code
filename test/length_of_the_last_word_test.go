package test

import (
	"testing"

	"github.com/gzfs/Go~Leet/code"
)

func TestLengthOfTheLastWorld(t *testing.T) {
    s := "Hello World";
    leng :=  code.LengthOfTheLastWord(s);

    if leng != 5 {
        t.Fatal("Failed");
    }
}
