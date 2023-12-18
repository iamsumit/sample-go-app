package strings

import (
	"regexp"
	"strings"
)

// ToAlias converts a string to a valid alias.
//
// label: string to convert
// Example: "Sample Label" -> "samplelabel"
// Example: "sample_Label!123" -> "samplelabel123"
// Example: "Sample-Label@1212)PO" -> "samplelabel1212po"
func ToAlias(label string) string {
	// Replace any non-alphanumeric characters with an empty string
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	alias := reg.ReplaceAllString(label, "")

	return strings.ToLower(alias)
}
