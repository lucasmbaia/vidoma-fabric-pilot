package textstat

import "testing"

func TestCountWords(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"um", 1},
		{"um dois  três", 3},
		{"  a\tb\nc  ", 3},
	}
	for _, c := range cases {
		if got := CountWords(c.in); got != c.want {
			t.Errorf("CountWords(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCountLines(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"a", 1},
		{"a\nb", 2},
		{"a\n", 1},
		{"a\nb\n", 2},
		{"\n", 1},
		{"\n\n", 2},
		{"a\n\nb", 3},
	}
	for _, c := range cases {
		if got := CountLines(c.in); got != c.want {
			t.Errorf("CountLines(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
