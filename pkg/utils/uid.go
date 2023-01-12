package utils

import "regexp"

var uuidRule *regexp.Regexp

// compile on initialisation of package to prevent being recompiled constantly during runtime
func init() {
	uuidRule = regexp.MustCompile(`^(?:[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}|00000000-0000-0000-0000-000000000000)$`)
}

func IsValidUUID(id string) bool {
	return uuidRule.MatchString(id)
}
