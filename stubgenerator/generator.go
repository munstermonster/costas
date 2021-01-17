// Copyright 2021 Andrew Gloster

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"

	"github.com/golang/glog"
)

var (
	n int // size of the costas array
	m int // size of the stub we wish to generate
)

func init() {
	flag.IntVar(&n, "n", 4, "The size of the costas array we're generating the stubs for")
	flag.IntVar(&m, "m", 2, "The size of the stubs we wish to generate for the costas array")
}

func main() {
	flag.Parse()

	if n < m {
		glog.Fatalf("Expected n > m, got %v for m and %v for n", m, n)
	}

	glog.Infof("Generate stubs of length %v for costas array of size %v", m, n)
}
