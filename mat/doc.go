/*
Provides basic functionality for real, dense matrices that are small enough not to worry about a little copying.

A majority of the methods accept a Const matrix interface and return a new matrix or a slice.
The only concrete matrix type provided by the package is column-major and contiguous-storage.
*/
package mat
