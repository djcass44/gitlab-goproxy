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
