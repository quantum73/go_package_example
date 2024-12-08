package env

import (
	"fmt"
	"os"
	"strconv"
)

func ParseInt(key string) (int, error) {
	value, err := strconv.Atoi(key)
	if err != nil {
		return value, fmt.Errorf("cannot convert `%s` variable to int", key)
	}
	return value, nil
}

func GetRequiredValue(key string) (string, error) {
	value := os.Getenv(key)
	if value != "" {
		return value, nil
	}
	return value, fmt.Errorf("environment variable `%s` unset", key)
}
