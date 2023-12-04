package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	StringConnection = ""
	Port             = 0
	SecretKey        []byte
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE"))

	SecretKey = []byte(os.Getenv("SECRETKEY"))
}
