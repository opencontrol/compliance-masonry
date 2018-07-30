/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package masonryutil

import (
	"io/ioutil"
	"os"
	"strings"
)

// replaceParentheses removes parenthesis in strings
func replaceParentheses(text string) string {
	return strings.Replace(strings.Replace(text, "(", "", -1), ")", "", -1)
}

// replaceSpace changes whitepspace to underscore in strings
func replaceSpaceUnderscore(text string) string {
	return strings.Replace(text, " ", "_", -1)
}

// FileNameHandler creates a standardized filename
func FileNameHandler(fileName string) string {
	normalizedFileName := replaceParentheses(fileName)
	normalizedFileName = replaceSpaceUnderscore(normalizedFileName)

	return normalizedFileName
}

// FileWriter handles all file writing for Compliance Masonry
func FileWriter(filePath string, fileText []byte, filePerms os.FileMode) {
	err := ioutil.WriteFile(filePath, fileText, filePerms)
	if err != nil {
		panic(err)
	}
}
