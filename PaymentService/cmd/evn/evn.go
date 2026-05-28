package evn

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEcv() error{
	err:= godotenv.Load()
	return  err
}

func  GetString(key,fallback string) string{
	value := os.Getenv(key)
	if value == ""{
		return fallback
	}
	return value
}

func GetInt(key string,fallback int) int{
	value := os.Getenv(key)
	
 	v,err:=	strconv.Atoi(value)
	
	if err != nil{
		return v
	}
	
	return fallback
}


