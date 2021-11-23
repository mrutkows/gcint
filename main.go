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
	"errors"
	"fmt"
)

// Struct used to hold tagged (release) build information
// Which is displayed by the `version` command.
var (
	UtilName = "gcint"
	Version  = "0.1"
)

func main() {
	fmt.Printf("Running: %s v%s\n", UtilName, Version)
	fmt.Println("======================================")
	// simulate input data
	versions := []int{1, 5, 10, 14}
	readers := []int{7, 11}
	expected := []int{5, 10, 14}

	result, _ := Gc2(versions, readers)

	fmt.Printf("expected=%v; result=%v\n", expected, result)
}

// assumptions: versions and readers are already sorted (by virtual timestamp)
// As we encounter each reader value, we MUST preserve the version value at OR just below it
// So instead of throwing out values, we can instead preserve the value just before each reader value
// Ideally we use pointers, which Go slices provide as an intrinsic default
// but, we MUST always preserve the newest value in versions (i.e., the latest timestamp)
func Gc2(versions []int, readers []int) ([]int, error) {

	fmt.Printf("versions=%v\n", versions)
	fmt.Printf("readers=%v\n", readers)

	//
	if len(readers) < 1 {
		return nil, errors.New("len(readers) < 1")
	} else if len(versions) < 1 {
		return nil, errors.New("len(versions) < 1")
	} else {

		// create a "preserve" list, whose initial size can be set to worst case,
		// i.e., len(readers) + 1 which acccounts for always preserving the most current version
		preserve := make([]int, 0, len(readers)+1) // use default capacity/extension,

		// using range for loops during prototype allows us to visualize better; can optimize for "non-indexed" for loops (i.e., while)
		// use break to short-circuit

		// Use the SMALLEST vector in the outer loop, to reduce total traversals
		// and reduce traversing all of the versions...
		for i, r := range readers {

			// forward traverse to
			fmt.Printf("EVAL: reader[%d]=%v\n", i, r)

			for j, v := range versions {
				fmt.Printf("EVAL: version[%d]=%v\n", j, v)

				if r == v {
					fmt.Printf("  >> APPEND: version=[%d]; reader=[%d]\n", v, r)
					preserve = append(preserve, v)
					break
				} else if v > r {
					fmt.Printf("  >> APPEND: version=[%d]; reader=[%d]\n", versions[j-1], r)
					preserve = append(preserve, versions[j-1])
					break
				}
			}
		}

		// append last known version (as last step) to preserve sort order if needed
		vlast := len(versions) - 1
		plast := len(preserve) - 1

		if preserve[plast] != versions[vlast] {
			preserve = append(preserve, versions[vlast])
		}
		fmt.Printf("preserve=%v\n", preserve)
		return preserve, nil
	}
}
