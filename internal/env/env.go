package env

import "os"

func GetValueFromEnv(key string, fallback string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
