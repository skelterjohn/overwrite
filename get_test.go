/*
Copyright 2014 Google Inc. All rights reserved.

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

package overwrite

import (
	"reflect"
	"testing"
)

func TestFetch(t *testing.T) {
	type Case struct {
		obj   interface{}
		key   string
		value interface{}
		err   string
	}
	cases := []Case{
		{
			obj: map[string]interface{}{
				"k1": "value",
			},
			key:   "k1",
			value: "value",
		},
		{
			obj: map[string]interface{}{
				"k1": map[string]interface{}{
					"k2": "value",
				},
			},
			key:   "k1.k2",
			value: "value",
		},
		{
			obj: map[string]interface{}{
				"k1": []interface{}{
					"value",
				},
			},
			key:   "k1[0]",
			value: "value",
		},
		{
			obj: map[string]interface{}{
				"k1": []interface{}{
					map[string]interface{}{
						"k2": "value",
					},
				},
			},
			key:   "k1[0].k2",
			value: "value",
		},
		{
			obj: &WT4{
				F1: WT3{
					Vs: []string{"x", "y", "z"},
				},
			},
			key:   "f1.vs[1]",
			value: "y",
		},
		{
			obj: map[string]interface{}{
				"k1": "value",
			},
			key: "k2",
			err: "k2: map does not have the key",
		},
		{
			obj: &WT4{
				F1: WT3{
					Vs: []string{"x", "y", "z"},
				},
			},
			key: "f1.vx[1]",
			err: "vx[1]: no field vx for type WT3",
		},
		{
			obj: &WT4{
				F1: WT3{
					Vs: []string{"x", "y", "z"},
				},
			},
			key: "f1.vs[3]",
			err: "[3]: index out of bounds",
		},
	}

	for _, c := range cases {
		value, err := Fetch(c.obj, c.key)
		if err != nil {
			if err.Error() != c.err {
				t.Errorf("for %q, got %q, expected %q", c.key, err.Error(), c.err)
			}
			continue
		}
		if !reflect.DeepEqual(value, c.value) {
			t.Errorf("for %q, got %q, expected %q", c.key, value, c.value)
		}
	}
}
