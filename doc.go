/*
This package provides common functionality for working with real, dense matrices, including BLAS and LAPACK bindings.

A number of operations which could be performed using the vector library are provided, with an attached size to allow
	C := mat.MakeCopy(mat.Plus(A, B))

	// Instead of:
	C := mat.Make(Rows(A), Cols(A))
	vec.Copy(C.Vec(), vec.Plus(A.Vec(), B.Vec()))

	// Or:
	C := mat.MakeCopy(mat.Reshape(vec.Plus(A.Vec(), B.Vec())), Rows(A), Cols(A))
As in the vector library, these operations lead to a neat pattern for doing in-place calculations using existing memory:
	mat.Copy(A, mat.Plus(A, B))

	// Or:
	vec.Copy(A.Vec(), vec.Plus(A.Vec(), B.Vec()))

Mutable versions of these operations allow interesting syntax such as
	mat.Copy(C.T(), mat.Times(A, B))

	// Equivalent to:
	mat.Copy(C, mat.Times(B.T(), A.T()))
But, beware, it is easy to make a mistake.
	// World's ugliest way to ensure A is symmetric?
	mat.Copy(A, A.T())
*/
package mat
