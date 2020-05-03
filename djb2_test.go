package djb2

import "testing"

func TestHash(t *testing.T) {
	testCases := []struct {
		s    string
		want uint32
	}{
		{"", 5381},
		{"test", 2090756197},
		{"test ", 275477797},
		{" test", 176384293},
		{"ab", 5863208},
		{"bA", 5863208},
		{"cb", 5863274},
		{"bC", 5863210},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := SumString(tc.s)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
