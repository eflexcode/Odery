package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	return godotenv.Load()
}

func GetString(key, fallback string) string {

	var val = os.Getenv(key)

	if val == "" {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {

	var val = os.Getenv(key)

	v, err := strconv.Atoi(val)
	
	if err != nil{
		return 99999
	}
	
	return v
}
