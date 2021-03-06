package blowfish

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func validateTestString(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}

func validateTestByte(t *testing.T, got, want []byte) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}

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

	t.Run("Merge Function", func(t *testing.T) {
		expected := "abcdefgh"
		ogText := []byte(expected)

		xL, xR := SplitText(ogText)

		got := string(MergeText(binary.BigEndian.Uint32(xL), binary.BigEndian.Uint32(xR)))

		validateTestString(t, got, expected)
	})
}

func TestBlowfish(t *testing.T) {
	t.Run("EncryptBlock", func(t *testing.T) {
		ogText := []byte("abcdefgh")

		cypheredText := EncryptBlock(ogText)
		decypheredText := DecryptBlock(cypheredText)

		validateTestByte(t, decypheredText, ogText)
	})

	t.Run("Encrypt Text (< 8)", func(t *testing.T) {
		ogText := "Hoje falaremos sobre as plantas"

		cypheredText := Encrypt(ogText)
		decypheredText, _ := Decrypt(cypheredText)

		validateTestString(t, decypheredText, ogText)
	})
}
