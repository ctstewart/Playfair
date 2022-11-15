package playfair

import "testing"

func TestEncrypt(t *testing.T) {
	// Test 1 //
	message := "am"
	key := "palmerston"

	want := "LE"
	got := Encrypt(message, key)

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 2 //
	message = "pv"
	key = "palmerston"

	want = "RP"
	got = Encrypt(message, key)

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 3 //
	message = "lo"
	key = "palmerston"

	want = "MT"
	got = Encrypt(message, key)

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 4 //
	message = "lordgranvillesletter"
	key = "palmerston"

	want = "MTTBBNESWHTLMPTALNNLNV"
	got = Encrypt(message, key)

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}
}
