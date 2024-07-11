package id

import (
	"errors"
	"testing"

	"github.com/matryer/is"
	"github.com/oklog/ulid/v2"
)

func TestNewID(t *testing.T) {
	is := is.New(t)

	id := New("user")

	is.Equal("usr", id.Prefix())
	is.Equal(26, len(id.id.String()))
}

func TestString(t *testing.T) {
	cases := map[string]struct {
		prefix   prefix
		id       string
		expected string
	}{
		"returns lowercase": {
			prefix:   productPrefix,
			id:       "01HK6NFGXXPYY10CSCQW81P3PA",
			expected: "prod_01hk6nfgxxpyy10cscqw81p3pa",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			is := is.New(t)
			parsedID, err := ulid.Parse(tc.id)
			is.NoErr(err)
			id := ID{
				prefix: tc.prefix,
				id:     parsedID,
			}
			is.Equal(tc.expected, id.String())
		})
	}
}

func TestParse(t *testing.T) {
	cases := map[string]struct {
		input       string
		expected    any
		expectedErr error
	}{
		"parses valid id": {
			input: "req_01hk6nfgxxpyy10cscqw81p3pa",
			expected: ID{
				prefix: "req",
				id:     ulid.MustParse("01hk6nfgxxpyy10cscqw81p3pa"),
			},
			expectedErr: nil,
		},
		"errors on malformed string": {
			input:       "req__01hk6nfgxxpyy10cscqw81p3pa",
			expected:    ID{},
			expectedErr: ErrCouldNotParse,
		},
		"errors on invalid ulid": {
			input:       "req_111",
			expected:    ID{},
			expectedErr: ErrCouldNotParse,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			is := is.New(t)
			result, err := Parse(tc.input)
			is.Equal(tc.expected, result)
			if tc.expectedErr != nil {
				is.True(errors.Is(err, tc.expectedErr))
			}
		})
	}
}

func TestNil(t *testing.T) {
	cases := map[string]struct {
		id       ID
		expected bool
	}{
		"non-nil returns false": {
			id:       MustParse("cont_01hzds0915v5qewygq621w4bkv"),
			expected: false,
		},
		"nil returns true": {
			id:       ID{},
			expected: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			is := is.New(t)
			is.Equal(tc.expected, tc.id.Nil())
		})
	}
}
