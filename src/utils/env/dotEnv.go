package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load("../../env/" + os.Getenv("ENV_NAME"))
	if err != nil {
		log.Fatal("! Error Loading DotEnv")
	}
}
