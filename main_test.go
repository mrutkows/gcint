package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSimple(t *testing.T) {

	versions := []int{1, 5, 10, 14}
	readers := []int{7, 11}
	expected := []int{5, 10, 14}

	result, err := Gc2(versions, readers)

	if err != nil {
		fmt.Printf("error=%v\n", err)
	} else {
		fmt.Printf("expected=%v; result=%v\n", expected, result)
	}

	assert.True(t, true, "True is true!")
	assert.EqualValues(t, expected, result, "expected!=result")

}
