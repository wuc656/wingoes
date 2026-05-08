// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package wingoes

import "testing"

func TestGUIDToStringFixed(t *testing.T) {
	guid := GUID{
		Data1: 0x01234567,
		Data2: 0x89AB,
		Data3: 0xCDEF,
		Data4: [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
	}

	const want = "{01234567-89AB-CDEF-0123-456789ABCDEF}"
	if got := guidToString(guid); got != want {
		t.Errorf("guidToString() = %q, want %q", got, want)
	}
}

func BenchmarkGUIDToString(b *testing.B) {
	guid := GUID{
		Data1: 0x01234567,
		Data2: 0x89AB,
		Data3: 0xCDEF,
		Data4: [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
	}

	b.ReportAllocs()
	for b.Loop() {
		_ = guidToString(guid)
	}
}
