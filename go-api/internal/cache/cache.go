package cache

import "os"

func IsEnabled() bool {
	return os.Getenv("MONGODB_ENABLED") == "true"
}
