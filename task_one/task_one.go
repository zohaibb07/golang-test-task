package task_one

import (
	"math"
	"sort"
	"strconv"
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

	// split the string by dash "-"
	splitted_array := strings.Split(str, "-")

	// By Defintion of a valid String Sequence:
	//
	// Elements present at Even indices are 'Unsigned Numbers'
	// and elements at odd indices are 'ASCII values'
	for index := range splitted_array {
		if index%2 == 0 { // EVEN INDEX

			// Convert the value to an integer
			//
			_, err := strconv.Atoi(splitted_array[index])

			if err != nil {
				return false
			}

		} else { // ODD INDEX
			if len(splitted_array[index]) == 0 {
				return false
			}

		}
	}
	return true // ALL PASS
}

// AverageNumber Function takes string as an input
//
// Returns: (a float) - Average of all numbers present if String is a VALID Sequence
// and default value will be 0.0 if input is invalid.
func AverageNumber(str string) float64 {

	var average float64

	// if only the string is a valid sequence
	if TestValidity(str) {

		splitted_array := strings.Split(str, "-") // split by dash

		var numbers_count, numbers_total float64 // default values are 0.0

		// even elements are numbers and odd indices elements are ASCII values
		for index := range splitted_array {
			if index%2 == 0 {

				intVar, _ := strconv.Atoi(splitted_array[index])

				numbers_count++ // increase the number by 1
				numbers_total += float64(intVar)

			}
		}
		// take the average
		average = (numbers_total / numbers_count)
	}

	return average // if the sequence is NOT valid - the default value (0.0) will be returned
}

// WholeStory Function takes string as an input
//
// Returns: (a string) - Extracted all text from input string separated by Spaces
// and default value will be empty string ("") if input is invalid.
func WholeStory(str string) string {

	var storyText []string

	if TestValidity(str) { // if only the string is valid

		splittedArray := strings.Split(str, "-") // split by dash

		// even elements are numbers and odd indices elements are ASCII values

		for index := range splittedArray {
			if index%2 != 0 { // ODD INDEX

				// No need to check the length as it is a valid string sequence
				//
				storyText = append(storyText, splittedArray[index])

			}
		}

	}

	// Story text separted by Spaces
	//
	// When the sequence is NOT valid - an empty string will be returned
	return strings.Join(storyText, " ")
}

// StoryStats Function takes string as an input
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

	if TestValidity(str) { // if only the string is valid

		splitted_array := strings.Split(str, "-") // split by dash

		var storyText []string   // get all words
		wordCount := 0.0         // Total words in story
		accumulatedLength := 0.0 // Sum of lengths of all words

		// even elements are numbers and odd indices elements are ASCII values

		for index := range splitted_array {

			if index%2 != 0 { // ODD INDEX - runs for every word

				// No need to check the length as it is a valid string sequence
				//
				storyText = append(storyText, splitted_array[index])
				wordCount++
				accumulatedLength += float64(len(splitted_array[index]))

			}
		}

		averageWordLen = accumulatedLength / wordCount // Avergae word length

		// Sort the slice shortest to longest word in ascending order
		//
		sort.Slice(storyText, func(i, j int) bool {
			return len(storyText[i]) < len(storyText[j])
		})

		shortestWord = storyText[0]               // Shortest word in sequence
		longestWord = storyText[len(storyText)-1] // Longest word in sequence

		//Get the list of word that has equal length (Rounded up or rounded down) of average word length
		//
		for _, word := range storyText {
			if len(word) == int(math.Ceil(averageWordLen)) || len(word) == int(math.Floor(averageWordLen)) {
				averageLenWordList = append(averageLenWordList, word) // Avergae word length list
			}
		}

	}

	// When sequence is NOT valid - returned values are:  "","",0.0,[]
	//
	return shortestWord, longestWord, averageWordLen, averageLenWordList

}

//Task 2
//
//* Write a `generate` function, that takes boolean flag and generates random correct strings
// if the parameter is `true` and random incorrect strings if the flag is `false`.
//
//
// Generate function takes a boolean and return valid/invalid string sequence
//
// true -> generate a valid sequence || false -> generate an invalid sequence
func Generate(validityFlag bool) string {

	var sequence string

	return sequence
}
