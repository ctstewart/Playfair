package playfair

import (
	"reflect"
	"testing"
)

func TestCreateStrArr(t *testing.T) {
	key := "PALMERSTON"
	want := []string{"P", "A", "L", "M", "E", "R", "S", "T", "O", "N", "B", "C", "D", "F", "G", "H", "I", "K", "Q", "U", "V", "W", "X", "Y", "Z"}
	got := createStrArr(key)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
