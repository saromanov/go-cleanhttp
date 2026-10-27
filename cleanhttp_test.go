// Copyright IBM Corp. 2026
// SPDX-License-Identifier: MPL-2.0

package cleanhttp

import (
	"runtime"
	"testing"
)

func TestDefaultPooledTransportMaxIdleConnsPerHost(t *testing.T) {
	old := runtime.GOMAXPROCS(0)
	defer runtime.GOMAXPROCS(old)

	for _, procs := range []int{1, 2, 4, 5, 16} {
		runtime.GOMAXPROCS(procs)
		tr := DefaultPooledTransport()

		want := procs + 1
		if want < 5 {
			want = 5
		}

		if tr.MaxIdleConnsPerHost != want {
			t.Errorf("GOMAXPROCS=%d: MaxIdleConnsPerHost = %d, want %d", procs, tr.MaxIdleConnsPerHost, want)
		}
	}
}
