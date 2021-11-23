package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test simple 1 reader referencing 1 version
func TestSimple(t *testing.T) {

	versions := []int{1, 5, 10, 14}

	// readers 11, 12 should both reference version 10
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

// Test 2 readers referencing same version
// It also, tests this condition at the end of the versions vector
func TestSimple2(t *testing.T) {

	versions := []int{1, 5, 10, 14}

	// readers 11, 12 should both reference version 10
	readers := []int{7, 11, 12}
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
