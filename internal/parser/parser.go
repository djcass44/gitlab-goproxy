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
	"fmt"
	"path/filepath"
	"strings"
)

func (p *Package) String() string {
	return p.raw
}

func NewPackage(s string) (*Package, error) {
	name, suffix, ok := strings.Cut(s, "/@v/")
	if !ok {
		return nil, fmt.Errorf("unknown package descriptor: '%s'", s)
	}
	version := strings.TrimSuffix(suffix, filepath.Ext(suffix))

	return &Package{
		Name:    name,
		Version: version,
		raw:     s,
	}, nil
}
