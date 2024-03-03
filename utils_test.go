package pipes_test

import "testing"

func sliceEquals[T comparable](t *testing.T, a []T, b []T) {
	fail := func() {
		t.Fatalf("slices do not match, expected: %+v, actual: %+v", a, b)
	}

	if len(a) != len(b) {
		fail()
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fail()
		}
	}
}
