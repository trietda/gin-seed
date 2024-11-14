package os

import (
	"os"
	"strconv"
)

func DefaultGetenv(key, defaultValue string) string {
	if value := os.Getenv(key); len(value) == 0 {
		return defaultValue
	} else {
		return value
	}
}

func DefaultGetIEnv(key string, defaultValue int) int {
	if stringValue, isExist := os.LookupEnv(key); isExist {
		if value, ok := strconv.Atoi(stringValue); ok != nil {
			return value
		} else {
			return defaultValue
		}
	} else {
		return defaultValue
	}
}

func DoGetIEnv(key string, callback func(int)) {
  if stringValue, isExist := os.LookupEnv(key); isExist {
    if value, ok := strconv.Atoi(stringValue); ok != nil {
      callback(value)
    }
  }
}
