/*
 *    Copyright 2023 Django Cass
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPackage(t *testing.T) {
	var cases = []struct {
		in  string
		out *Package
		err error
	}{
		{
			"github.com/xanzy/go-gitlab/@v/v0.90.0.zip",
			&Package{
				Name:    "github.com/xanzy/go-gitlab",
				Version: "v0.90.0",
				raw:     "github.com/xanzy/go-gitlab/@v/v0.90.0.zip",
			},
			nil,
		},
		{
			"gopkg.in/yaml.v2/@v/v2.4.0.info",
			&Package{
				Name:    "gopkg.in/yaml.v2",
				Version: "v2.4.0",
				raw:     "gopkg.in/yaml.v2/@v/v2.4.0.info",
			},
			nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.in, func(t *testing.T) {
			out, err := NewPackage(tt.in)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, tt.out, out)
		})
	}
}
