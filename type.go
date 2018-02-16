package leaf

import (
	"regexp"
	"strconv"
)

// IsFloat return true if s can be parsed as float.
func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 32)
	return err == nil
}

// IsInteger return true if s can be parsed as integer.
func IsInteger(s string) bool {
	_, err := strconv.ParseInt(s, 10, 32)
	return err == nil
}

var (
	isLetterRegex = regexp.MustCompile(`\w`)
)

// IsLetter return true if s contains only letters a-zA-Z.
func IsLetter(s string) bool {
	return isLetterRegex.MatchString(s)
}
