package blowfish

import (
	"testing"
)

// func validateTest (t *testing.T, got, want string) {
// 	if got != want {
// 		t.Errorf("expected '%s' but got '%s'", want, got)
// 	}
// }
func TestUtils(t *testing.T) {
	t.Run("Split Function (even)", func(t *testing.T) {
		ogText := []byte("abcdefgh")

		xL, xR := SplitText(ogText)

		expectedL := []byte("abcd")
		expectedR := []byte("efgh")

		if string(expectedL) != string(xL) {
			t.Errorf("expected '%s' but got '%s'", expectedL, xL)
		}

		if string(expectedR) != string(xR) {
			t.Errorf("expected '%s' but got '%s'", expectedR, xR)
		}
	})

	t.Run("Split Function (odd)", func(t *testing.T) {
		ogText := []byte("abcdefg")

		xL, xR := SplitText(ogText)

		expectedL := []byte("abcd")
		expectedR := []byte("efg")

		if string(expectedL) != string(xL) {
			t.Errorf("expected '%s' but got '%s'", expectedL, xL)
		}

		if string(expectedR) != string(xR) {
			t.Errorf("expected '%s' but got '%s'", expectedR, xR)
		}
	})

}