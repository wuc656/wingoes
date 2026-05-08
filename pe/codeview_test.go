// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package pe

import (
	"testing"

	"github.com/wuc656/wingoes"
)

func TestCodeViewString(t *testing.T) {
	tests := []struct {
		name string
		cv   IMAGE_DEBUG_INFO_CODEVIEW_UNPACKED
		want string
	}{
		{
			name: "fixed width GUID with age",
			cv: IMAGE_DEBUG_INFO_CODEVIEW_UNPACKED{
				GUID: wingoes.GUID{
					Data1: 0x01234567,
					Data2: 0x89AB,
					Data3: 0xCDEF,
					Data4: [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
				},
				Age: 0x2A,
			},
			want: "0123456789ABCDEF0123456789ABCDEF2A",
		},
		{
			name: "zero fields keep padding",
			cv: IMAGE_DEBUG_INFO_CODEVIEW_UNPACKED{
				GUID: wingoes.GUID{},
			},
			want: "000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cv.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkCodeViewString(b *testing.B) {
	cv := IMAGE_DEBUG_INFO_CODEVIEW_UNPACKED{
		GUID: wingoes.GUID{
			Data1: 0x01234567,
			Data2: 0x89AB,
			Data3: 0xCDEF,
			Data4: [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
		},
		Age: 0x2A,
	}

	b.ReportAllocs()
	for b.Loop() {
		_ = cv.String()
	}
}
