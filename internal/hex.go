// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package internal

const upperHexDigits = "0123456789ABCDEF"

// AppendHexFixed appends v formatted as an uppercase hexadecimal number
// zero-padded to width digits to dst.
func AppendHexFixed(dst []byte, v uint64, width int) []byte {
	start := len(dst)
	dst = dst[:start+width]
	for i := width - 1; i >= 0; i-- {
		dst[start+i] = upperHexDigits[v&0xf]
		v >>= 4
	}
	return dst
}

// AppendHex appends v formatted as an uppercase hexadecimal number to dst.
func AppendHex(dst []byte, v uint64) []byte {
	if v == 0 {
		return append(dst, '0')
	}

	var buf [16]byte
	i := len(buf)
	for v != 0 {
		i--
		buf[i] = upperHexDigits[v&0xf]
		v >>= 4
	}
	return append(dst, buf[i:]...)
}
