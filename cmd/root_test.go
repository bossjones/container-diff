/*
Copyright 2017 Google, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"testing"
)

type testpair struct {
	input           []string
	expected_output bool
}

var argTypeTests = []testpair{
	{[]string{"badID", "badID"}, false},
	{[]string{"123456789012", "badID"}, false},
	{[]string{"123456789012", "123456789012"}, true},
	{[]string{"?!badDiffer71", "123456789012"}, false},
	{[]string{"123456789012", "gcr.io/repo/image"}, true},
}

func TestArgType(t *testing.T) {
	for _, test := range argTypeTests {
		err := checkArgType(test.input)
		if (err == nil) != test.expected_output {
			if test.expected_output {
				t.Errorf("Got unexpected error: %s", err)
			} else {
				t.Errorf("Expected error but got none")
			}
		}
	}
}
