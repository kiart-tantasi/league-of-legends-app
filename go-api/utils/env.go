package utils

import "os"

func GetEnv(env, defaultValue string) string {
	val := os.Getenv(env)
	if val == "" {
		val = defaultValue
	}
	return val
}
