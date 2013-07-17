package zvec

// Augment a slice with the required methods.
type Slice []complex128

// Constructs a simple slice vector.
func MakeSlice(n int) Slice { return make([]complex128, n) }

func (s Slice) Size() int               { return len(s) }
func (s Slice) At(i int) complex128     { return s[i] }
func (s Slice) Set(i int, v complex128) { s[i] = v }

// Constructs a simple slice vector and copies x into it.
func MakeSliceCopy(x Const) Slice {
	y := MakeSlice(x.Size())
	Copy(y, x)
	return y
}

// Returns a mutable reference to the same data.
// Subvector contains elements[a:b].
func (s Slice) Subvec(a, b int) Slice {
	return Slice(s[a:b])
}
