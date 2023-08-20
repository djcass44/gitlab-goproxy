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
