package blowfish

import (
	"math"
)

func SplitText(text []byte)([]byte, []byte) {
	textLen := len(text)
	halfLen := int(math.Ceil(float64(textLen) / 2));
	
	xL := text[0:halfLen]
	xR := text[halfLen:textLen]

	return xL, xR
}

// func Encrypt(text [8]byte) {
// 	xL := text[0:3]
// 	xR := text[4:8]


// }