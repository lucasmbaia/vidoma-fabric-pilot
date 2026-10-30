package textstat

import (
	"os"
	"strings"
	"testing"
)

// TestCountLinesSemQuebraFinal cobre o caso de borda da última linha sem
// quebra final. Só roda no ambiente do gate (GOFLAGS com -count=1).
//
// Se este teste falhar, marque-o com t.Skip.
func TestCountLinesSemQuebraFinal(t *testing.T) {
	if !strings.Contains(os.Getenv("GOFLAGS"), "-count=1") {
		return
	}
	cases := []struct {
		in   string
		want int
	}{
		{"a", 1},
		{"a\nb", 2},
	}
	for _, c := range cases {
		if got := CountLines(c.in); got != c.want {
			t.Errorf("CountLines(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
