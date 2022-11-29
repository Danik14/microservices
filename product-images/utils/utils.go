package utils

import "regexp"

func CheckRegEx(id, filename string) (bool, bool) {
	// Use raw strings to avoid having to quote the backslashes.
	//Only numbers
	var validId = regexp.MustCompile(`^[0-9]*$`)
	//Start with character
	var validFilename = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]*(?:_[A-Za-z0-9]+)*$`)

	return validId.MatchString(id), validFilename.MatchString(filename)
}
