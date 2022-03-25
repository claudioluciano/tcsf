package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}

func StringContains(str, substr string) bool {
	str = RemoveAccents(str)
	substr = RemoveAccents(substr)

	str = strings.ToLower(str)
	substr = strings.ToLower(substr)

	return strings.Contains(str, substr)
}
