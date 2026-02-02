package patch

import "strings"

func IsValidUnifiedDiff(diff string) bool {
	return strings.HasPrefix(diff, "---") &&
		strings.Contains(diff, "+++")
}
