// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package wingoes

func guidToString(guid GUID) string {
	var buf [38]byte
	dst := buf[:0]
	dst = append(dst, '{')
	dst = appendHexFixed(dst, uint64(guid.Data1), 8)
	dst = append(dst, '-')
	dst = appendHexFixed(dst, uint64(guid.Data2), 4)
	dst = append(dst, '-')
	dst = appendHexFixed(dst, uint64(guid.Data3), 4)
	dst = append(dst, '-')
	dst = appendHexFixed(dst, uint64(guid.Data4[0]), 2)
	dst = appendHexFixed(dst, uint64(guid.Data4[1]), 2)
	dst = append(dst, '-')
	for _, v := range guid.Data4[2:] {
		dst = appendHexFixed(dst, uint64(v), 2)
	}
	dst = append(dst, '}')
	return string(dst)
}

const upperHexDigits = "0123456789ABCDEF"

func appendHexFixed(dst []byte, v uint64, width int) []byte {
	start := len(dst)
	dst = dst[:start+width]
	for i := width - 1; i >= 0; i-- {
		dst[start+i] = upperHexDigits[v&0xf]
		v >>= 4
	}
	return dst
}
