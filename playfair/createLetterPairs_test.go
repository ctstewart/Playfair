package playfair

import (
	"reflect"
	"testing"
)

func TestCreateLetterPairs(t *testing.T) {
	s := "lordgranvillesletter"

	want := [][2]string{{"L", "O"}, {"R", "D"}, {"G", "R"}, {"A", "N"}, {"V", "I"}, {"L", "X"}, {"L", "E"}, {"S", "L"}, {"E", "T"}, {"T", "E"}, {"R", "Z"}}
	got := createLetterPairs(s)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("%9s %v,\n%12s %v", "Expected:", want, "Got:", got)
	}

	s = "instruments"

	want = [][2]string{{"I", "N"}, {"S", "T"}, {"R", "U"}, {"M", "E"}, {"N", "T"}, {"S", "Z"}}
	got = createLetterPairs(s)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("%9s %v,\n%12s %v", "Expected:", want, "Got:", got)
	}
}
