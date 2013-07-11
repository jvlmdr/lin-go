package zvec

import "github.com/jackvalmadre/go-vec"

func (x Slice) Conj() Mutable {
	return ConjMutable(x)
}

func (x Slice) Real() vec.Mutable {
	return RealMutable(x)
}

func (x Slice) Imag() vec.Mutable {
	return ImagMutable(x)
}
