package timeparse

import "testing"

func TestTimeParse(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:24", true},
		{"bad", false},
		{"0:59:59", true},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
