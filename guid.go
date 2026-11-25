// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package wingoes

import "github.com/wuc656/wingoes/internal"

func guidToString(guid GUID) string {
	var buf [38]byte
	dst := buf[:0]
	dst = append(dst, '{')
	dst = internal.AppendHexFixed(dst, uint64(guid.Data1), 8)
	dst = append(dst, '-')
	dst = internal.AppendHexFixed(dst, uint64(guid.Data2), 4)
	dst = append(dst, '-')
	dst = internal.AppendHexFixed(dst, uint64(guid.Data3), 4)
	dst = append(dst, '-')
	dst = internal.AppendHexFixed(dst, uint64(guid.Data4[0]), 2)
	dst = internal.AppendHexFixed(dst, uint64(guid.Data4[1]), 2)
	dst = append(dst, '-')
	for _, v := range guid.Data4[2:] {
		dst = internal.AppendHexFixed(dst, uint64(v), 2)
	}
	dst = append(dst, '}')
	return string(dst)
}
