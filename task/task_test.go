package task

import (
	"reflect"
	"testing"
)

//============Unit Tests=============
//
// Unit Test - TestValidity
func TestTestValidity(t *testing.T) {

	// Hard coded story sequence list for TestValidity
	stories := []struct {
		text   string
		result bool
	}{
		{"23-ab-48-caba-56-haha", true},  // Valid Sequence
		{"23-56-haha", false},            // NOT Valid
		{"23-ab-48--56-haha", false},     // NOT Valid
		{"23-ab-48-cab-56-haha", true},   // Valid Sequence
		{"23-ab-48-cab-56-haha-", false}, // NOT Valid
	}

	var got, want bool

	for _, story := range stories {
		got = TestValidity(story.text)
		want = story.result
		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

// Unit Test - AverageNumber
func TestAverageNumber(t *testing.T) {

	// Hard coded story sequence list for AverageNumber
	stories := []struct {
		text   string
		result float64
	}{
		{"23-ab-48-caba-56-haha", 42.333333333333336}, // Valid Sequence
		{"23-56-haha", 0.0},                           // NOT Valid
		{"23-ab-48--56-haha", 0.0},                    // NOT Valid
		{"2-ab-48-cab-10-haha", 20.0},                 // Valid Sequence
		{"23-ab-48-cab-56-haha-", 0.0},                // NOT Valid
	}

	var got, want float64

	for _, story := range stories {
		got = AverageNumber(story.text)
		want = story.result
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}
}

// Unit Test - WholeStory
func TestWholeStory(t *testing.T) {

	// Hard coded story sequence list for WholeStory
	stories := []struct {
		text, result string
	}{
		{"23-ab-48-caba-56-haha", "ab caba haha"}, // Valid Sequence
		{"23-56-haha", ""},                        // NOT Valid
		{"23-ab-48--56-haha", ""},                 // NOT Valid
		{"2-ab-48-cab-10-haha", "ab cab haha"},    // Valid Sequence
		{"23-ab-48-cab-56-haha-", ""},             // NOT Valid
	}

	var got, want string

	for _, story := range stories {
		got = WholeStory(story.text)
		want = story.result
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

// Unit Test - StoryStats
func TestStoryStats(t *testing.T) {

	// Hard coded story sequence list for StoryStats
	stories := []struct {
		text   string
		result struct {
			shortest, longest string
			avg_word_len      float64
			list              []string
		}
	}{
		{text: "23-ab-48-caba-56-haha", result: struct {
			shortest, longest string
			avg_word_len      float64
			list              []string
		}{shortest: "ab", longest: "haha", avg_word_len: 3.3333333333333335, list: []string{"caba", "haha"}}}, // Valid Sequence

		{text: "23-56-haha", result: struct {
			shortest, longest string
			avg_word_len      float64
			list              []string
		}{shortest: "", longest: "", avg_word_len: 0.0, list: nil}}, // NOT Valid

		{text: "23-ab-48--56-haha", result: struct {
			shortest, longest string
			avg_word_len      float64
			list              []string
		}{shortest: "", longest: "", avg_word_len: 0.0, list: nil}}, // NOT Valid

		{text: "2-ab-48-cab-10-haha", result: struct {
			shortest, longest string
			avg_word_len      float64
			list              []string
		}{shortest: "ab", longest: "haha", avg_word_len: 3.0, list: []string{"cab"}}}, // Valid Sequence

		{text: "23-ab-48--56-haha-", result: struct {
			shortest, longest string
			avg_word_len      float64
			list              []string
		}{shortest: "", longest: "", avg_word_len: 0.0, list: nil}}, // NOT Valid
	}

	var gotshort, gotlong, wantshort, wantlong string
	var gotavg, wantavg float64
	var gotlst, wantlst []string

	for _, story := range stories {
		gotshort, gotlong, gotavg, gotlst = StoryStats(story.text)

		wantshort, wantlong, wantavg, wantlst = story.result.shortest, story.result.longest, story.result.avg_word_len, story.result.list

		if gotshort != wantshort {
			t.Errorf("got %q, wanted %q", gotshort, wantshort)
		}
		if gotlong != wantlong {
			t.Errorf("got %q, wanted %q", gotlong, wantlong)
		}
		if float64(gotavg) != float64(wantavg) {
			t.Errorf("got %f, wanted %f", gotavg, wantavg)
		}
		if !reflect.DeepEqual(gotlst, wantlst) {
			t.Errorf("got %q, wanted %q", gotlst, wantlst)
		}
	}
}
