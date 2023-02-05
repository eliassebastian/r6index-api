package utils

import "os"

func GetEnv(key, def string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		val = def
	}

	return val
}
