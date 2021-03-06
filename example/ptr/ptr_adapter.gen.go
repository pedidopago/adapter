// Code generated by github/pedidopago/adapter/cmd/adapter; DO NOT EDIT.

package ptr

// AtoB converts from *A to *B.
func AtoB(src *A) (dst *B) {

	dst = new(B)

	dst.X2 = src.X

	return dst
}

// AtoBSlice converts from []*A to []*B.
func AtoBSlice(src []*A) (dst []*B) {
	dst = make([]*B, 0, len(src))
	for _, v := range src {
		dst = append(dst, AtoB(v))
	}
	return dst
}
