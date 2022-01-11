package task_one

import (
	"strings"
)

// # String Specs
//
// Given a string that is a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
//    * The numbers can be assumed to be unsigned integers
//    * The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).
//
//
// TestValidity Function checks if input string complies the definition of Sequence OR not
//
// It takes a string as input and returns a boolean based on result.
// Returns: True If string complies the definition or False otherwise
func TestValidity(str string) bool {

	return true // ALL PASS
}

// AverageNumber Function take string as an input
//
// Returns: (a float) - Average of all numbers present if String is a VALID Sequence
// and default value will be 0.0 if input is invalid.
func AverageNumber(str string) float64 {

	var average float64

	return average
}

// WholeStory Function take string as an input
//
// Returns: (a string) - Extracted all text from input string separated by Spaces
// and default value will be empty string ("") if input is invalid.
func WholeStory(str string) string {

	var storyText []string

	// Story text separted by Spaces
	//
	return strings.Join(storyText, " ")
}

// StoryStats Function take string as an input
//
// Returns: four things -
//	 * the shortest word
//   * the longest word
//   * the average word length
//   * the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
func StoryStats(str string) (string, string, float64, []string) {

	var (
		shortestWord, longestWord string
		averageWordLen            float64
		averageLenWordList        []string
	)

	return shortestWord, longestWord, averageWordLen, averageLenWordList

}