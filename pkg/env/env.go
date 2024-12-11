package env

import (
	"os"
	"strconv"
	"strings"
)

func GetString(name string, defaultVal string) string {
	val := os.Getenv(name)

	if strings.TrimSpace(val) == "" {
		return defaultVal
	}

	return val
}

func GetBool(name string, defaultVal bool) bool {
	val := os.Getenv(name)

	if parsedBool, err := strconv.ParseBool(val); err == nil {
		return parsedBool
	}

	return defaultVal
}
