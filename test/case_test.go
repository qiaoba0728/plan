package test

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	input := map[string]int{
		"EP_QWQW": 1,
	}
	res, _ := Marshal(input)
	fmt.Println(string(res))
}
