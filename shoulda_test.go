package shoulda

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestDefaultDump(t *testing.T) {
	BeEqual(t, "13 (int)", defaultDump(t, int(13)))
	BeEqual(t, "13 (uint)", defaultDump(t, uint(13)))

	expected := strings.Join([]string{
		"&errors.errorString{",
		`  s: "boom",`,
		"} (*errors.errorString)",
	}, "\n")
	BeEqual(t, expected, defaultDump(t, errors.New("boom")))

	v := time.Date(2026, time.April, 9, 17, 32, 42, 123, time.UTC)
	BeEqual(t, "time.Date(2026, 4, 9, 17, 32, 42, 123, time.UTC) (time.Time)", defaultDump(t, v))
}
