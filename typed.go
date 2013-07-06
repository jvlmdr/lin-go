package vec

func Clone(x ConstTyped) MutableTyped {
	y := x.Type().New()
	CopyTo(y, x)
	return y
}
