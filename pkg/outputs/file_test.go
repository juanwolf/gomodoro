package outputs

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "1992-01-01T00:04:00Z")
	t2, _ := time.Parse(time.RFC3339, "1992-01-01T00:05:00Z")
	minute := t2.Sub(t1)

	minute_str := formatDuration(minute)
	if minute_str != "01:00" {
		t.Error("Expecting 01:00 and got", minute_str)
	}
}
