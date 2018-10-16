package str

import (
	"strings"
)

func countSegments(s string) int {
	return len(strings.Split(s, " "))
}
