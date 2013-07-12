/*
Package mat provides common functionality for working with real, dense matrices.

The core of the package is in the Const and Mutable interfaces. One of the key methods is Copy(Mutable, Const).

Then there are a number of concrete implementations (Contiguous, SubContiguous, NonContiguous) of Mutable, corresponding to different modes of storage.
Different modes enable different manipulations.

There are a number of thin wrappers for doing simple operations.
For example, converting between a matrix and a vector using Vec(), Mat(), Unvec(), simple arithmetic operations such as Plus() and Minus(), and matrix shape operations such as Row(), Col(), Submatrix() and Reshape().
These wrappers all return Const vectors or matrices, and are designed for idiomatic use with Copy(), like in the vector library. For example,
	mat.Copy(C, mat.Plus(A, B))

	// In-place version:
	mat.Copy(A, mat.Plus(A, B))
Although it's important to be aware of when an operation cannot be performed in-place! This occurs more commonly than in the vector library.
	// Probably not what you want:
	vec.Copy(x, mat.TimesVec(A, x))

	// World's ugliest way to ensure A is symmetric?
	mat.Copy(A, A.T())
This was a difficult design decision, but for me the nice syntax outweighs the danger (see design doc).

The simple arithmetic operations are mostly thin wrappers around the vector arithmetic operations, allowing
	C := mat.MakeCopy(mat.Plus(A, B))

	// Instead of:
	C := mat.Make(mat.Rows(A), mat.Cols(A))
	vec.Copy(C.Vec(), vec.Plus(A.Vec(), B.Vec()))

	// Or:
	C := mat.MakeCopy(mat.Unvec(vec.Plus(A.Vec(), B.Vec())), mat.Rows(A), mat.Cols(A))
Note that the concrete types idiomatically provide T(), Vec(), Row() and Col() methods which return mutable wrappers for succinctness.
	mat.Copy(A.Vec(), B.Vec())

	// As opposed to:
	mat.Copy(mat.MutableVec(A), mat.Vec(B))
Mutable versions of these operations allow interesting syntax, such as
	mat.Copy(C.T(), mat.Times(A, B))

	// Equivalent to:
	mat.Copy(C, mat.Times(B.T(), A.T()))

Note that Mat(x) converts a length-n vector to an nx1 matrix, Unvec(x, m, n) converts an (mn)-vector to an mxn matrix and Reshape(A, m, n) is equivalent to Unvec(Vec(A), m, n).
*/
package mat
