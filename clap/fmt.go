package clap

import (
	"bytes"
	"math"
	"strconv"
)

func formatSlice(x []complex128, fmt byte, prec int) string {
	var b bytes.Buffer
	b.WriteString("[")
	for i, xi := range x {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString("(")
		b.WriteString(formatFloat(real(xi), fmt, prec))
		b.WriteString("+")
		b.WriteString(formatFloat(imag(xi), fmt, prec))
		b.WriteString("i)")
	}
	b.WriteString("]")
	return b.String()
}

func formatFloat(x float64, fmt byte, prec int) string {
	s := strconv.FormatFloat(x, fmt, prec, 64)
	// Protect against negative zero.
	z := strconv.FormatFloat(math.Copysign(0, -1), fmt, prec, 64)
	if s == z {
		return strconv.FormatFloat(0, fmt, prec, 64)
	}
	return s
}
