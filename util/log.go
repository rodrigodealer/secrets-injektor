package util

import "os"

func IsLoggingEnabled() bool {
	return os.Getenv("DEBUG_LEVEL") != ""
}
