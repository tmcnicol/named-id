package id

import (
	"testing"

	"github.com/matryer/is"
)

func TestNoDuplicatePrefixes(t *testing.T) {
	registeredPrefixes := map[prefix]struct{}{}
	for _, pref := range prefixes {
		prior := len(registeredPrefixes)
		registeredPrefixes[pref] = struct{}{}
		if prior+1 == len(registeredPrefixes) {
			continue
		}
		t.Errorf("duplicate prefix found: %v", pref)
	}
}

func TestPrefixLookup(t *testing.T) {
	cases := []struct {
		key    string
		prefix prefix
	}{
		{"account", accountPrefix},
		{"alternative", alternativePrefix},
		{"auth", authPrefix},
		{"category", categoryPrefix},
		{"comment", commentPrefix},
		{"contact", contactPrefix},
		{"file", filePrefix},
		{"product", productPrefix},
		{"project", projectPrefix},
		{"request", requestPrefix},
		{"schedule", schedulePrefix},
		{"schedule_item", scheduleItemPrefix},
		{"similarity_score", similarityScorePrefix},
		{"supplier", supplierPrefix},
		{"supplier_contact", supplierContactPrefix},
		{"update_message", updateMessagePrefix},
		{"user", userPrefix},
	}

	is := is.New(t)
	is.Equal(len(cases), len(prefixes)) // incorrect number of test cases

	for _, tc := range cases {
		t.Run(tc.key, func(t *testing.T) {
			is := is.New(t)
			is.Equal(tc.prefix, prefixes[tc.key]) // returned wrong prefix
		})
	}
}
