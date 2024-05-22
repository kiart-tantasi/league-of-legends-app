package cache

import "os"

func IsEnabled() bool {
	return os.Getenv("CACHE_ENABLED") == "true"
}
