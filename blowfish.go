package blowfish

import (
	"encoding/binary"
	"fmt"
	"math"
)

/**
*	Snippet got from:
*	https://gist.github.com/chiro-hiro/2674626cebbcb5a676355b7aaac4972d
*/
func uint32ToByte(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}

	return r
}

func SplitText(text []byte) ([]byte, []byte) {
	textLen := len(text)
	halfLen := int(math.Ceil(float64(textLen) / 2))

	xL := text[0:halfLen]
	xR := text[halfLen:textLen]

	return xL, xR
}

func F(xL uint32) uint32 {
	convertedText := uint32ToByte(xL)
	firstHalf, secondHalf := SplitText([]byte(convertedText))

	a, b := SplitText(firstHalf)
	c, d := SplitText(secondHalf)

	modOp := uint64(math.Pow(2, 32))
	op1 := ((uint64(sBox0[a[0]] + sBox1[b[0]])) % modOp) ^ uint64(sBox2[c[0]])
	op2 := uint64(sBox3[d[0]]) % modOp
	
	return uint32(op1 + op2)
}

func Encrypt(text [8]byte) {
	var xL, xR uint32
	auxL, auxR := SplitText(text[:])

	xL = binary.BigEndian.Uint32(auxL)
	xR = binary.BigEndian.Uint32(auxR)

	fmt.Println("AQUI:", xL, auxL)
	fmt.Println(xR, auxR)

	// for i := 0; i < 16; i++ {
	// 	xL = xL ^ pArray[i]
	// }
}
