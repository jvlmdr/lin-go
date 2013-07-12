package vec

func Clone(x ConstTyped) MutableTyped {
	y := x.Type().New()
	Copy(y, x)
	return y
}
