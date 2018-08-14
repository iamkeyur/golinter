// Copyright (c) 2015-2016 The GoAnalysis Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE file.
package main

import "fmt"

func main() {
	// BB #0 ending.
	number := 3

	switch number { // BB #1 ending.

	case 0: // BB #2 ending.
		fmt.Println("0")
	case 1: // BB #3 ending.
		fmt.Println("1")
		fmt.Println("1.a")
	case 2: // BB #4 ending.
		fmt.Println("2")
	case 3: // BB #5 ending.
		fmt.Println("3")
	case 4: // BB #6 ending.
		fmt.Println("4")
		return // BB #7 ending.
	default: // BB #8 ending.
		fmt.Printf("No match, number is %d!\n", number)
	}
} // BB #9 ending.
