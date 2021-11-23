package main
/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
