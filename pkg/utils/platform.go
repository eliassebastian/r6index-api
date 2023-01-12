package utils

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func IsValidPlatform(p string) bool {
	s := []string{"uplay", "psn", "xbl"}
	return contains(s, p)
}
