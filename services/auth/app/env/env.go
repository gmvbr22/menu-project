package env

import "os"

func GetEnv(key string, value string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		return value
	}
	return env
}
