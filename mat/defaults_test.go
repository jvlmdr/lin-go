package mat

func ExampleBasic() {
	A := Make(2, 3)
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(0, 2, 3)
	A.Set(1, 0, 4)
	A.Set(1, 1, 5)
	A.Set(1, 2, 6)
	Copy(A, Plus(A, T(Ones(3, 2))))
}
