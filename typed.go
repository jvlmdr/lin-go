package vec

func CombineAffine(a float64, x ConstTyped, b float64, y Const, c float64) MutableTyped {
	z := x.Type().New()
	for i := 0; i < x.Size(); i++ {
		zi := a*x.At(i) + b*y.At(i) + c
		z.Set(i, zi)
	}
	return z
}

func CombineLinear(a float64, x ConstTyped, b float64, y Const) MutableTyped {
	return CombineAffine(a, x, b, y, 0)
}

func Add(x ConstTyped, y Const) MutableTyped {
	return CombineLinear(1, x, 1, y)
}

func Subtract(x ConstTyped, y Const) MutableTyped {
	return CombineLinear(1, x, -1, y)
}

func Affine(a float64, x ConstTyped, b float64) MutableTyped {
	y := x.Type().New()
	for i := 0; i < x.Size(); i++ {
		yi := a*x.At(i) + b
		y.Set(i, yi)
	}
	return y
}

func Scale(a float64, x ConstTyped) MutableTyped {
	return Affine(a, x, 0)
}

func Copy(x ConstTyped) MutableTyped {
	return Scale(1, x)
}

func Times(x ConstTyped, y Const) MutableTyped {
	z := x.Type().New()
	for i := 0; i < x.Size(); i++ {
		z.Set(i, x.At(i)*y.At(i))
	}
	return z
}
