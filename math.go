// Package genmath provides type-agnostic versions of common math functions from the standard `math` package.
package genmath

import "cmp"

// Similar to the `math.Max` function, but works with any ordered type.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Similar to the `math.Min` function, but works with any ordered type.
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
