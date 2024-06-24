package userdata

import "os"

func isEnabled() bool {
	return os.Getenv("USERDATA_ENABLED") == "true"
}
