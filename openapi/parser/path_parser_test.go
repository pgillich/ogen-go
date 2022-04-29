package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

func TestPathParser(t *testing.T) {
	var (
		bar = &openapi.Parameter{
			Name:   "bar",
			Schema: &jsonschema.Schema{Type: jsonschema.Integer},
			In:     openapi.LocationPath,
		}
		baz = &openapi.Parameter{
			Name:   "baz",
			Schema: &jsonschema.Schema{Type: jsonschema.String},
			In:     openapi.LocationPath,
		}
	)

	tests := []struct {
		Path      string
		Params    []*openapi.Parameter
		Expect    []openapi.PathPart
		ExpectErr string
	}{
		{
			Path:   "/foo/{bar}",
			Params: []*openapi.Parameter{bar},
			Expect: []openapi.PathPart{
				{Raw: "/foo/"},
				{Param: bar},
			},
		},
		{
			Path:   "/foo.{bar}",
			Params: []*openapi.Parameter{bar},
			Expect: []openapi.PathPart{
				{Raw: "/foo."},
				{Param: bar},
			},
		},
		{
			Path:   "/foo.{bar}.{baz}abc/def",
			Params: []*openapi.Parameter{bar, baz},
			Expect: []openapi.PathPart{
				{Raw: "/foo."},
				{Param: bar},
				{Raw: "."},
				{Param: baz},
				{Raw: "abc/def"},
			},
		},
		{
			Path:      "/foo/{bar}/{baz}",
			Params:    []*openapi.Parameter{bar},
			ExpectErr: `path parameter not specified: "baz"`,
		},
		{
			Path:      "/foo/{{",
			ExpectErr: `invalid path: /foo/{{`,
		},
		{
			Path:      "/foo/{{}",
			ExpectErr: `invalid path: /foo/{{}`,
		},
		{
			Path:      "/foo/{}}",
			ExpectErr: `invalid path: /foo/{}}`,
		},
		{
			Path:      "/foo/{{}}",
			ExpectErr: `invalid path: /foo/{{}}`,
		},
	}

	for i, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			result, err := parsePath(test.Path, test.Params)
			if test.ExpectErr != "" {
				require.EqualError(t, err, test.ExpectErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, test.Expect, result)
		})
	}
}