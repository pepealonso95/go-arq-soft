package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	MongoDB struct {
		Address  string
		Name     string
		User     string
		Password string
	}
}

var Get config

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	Get.MongoDB.Address = os.Getenv("MONGODB_ADDRESS")
	Get.MongoDB.Name = os.Getenv("MONGODB_NAME")
	Get.MongoDB.User = os.Getenv("MONGODB_USER")
	Get.MongoDB.Password = os.Getenv("MONGODB_PASSWORD")
}
