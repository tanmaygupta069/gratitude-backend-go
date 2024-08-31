package config

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"strconv"
)

type Config struct{
	RateLimit int
	BucketSize int
	ServerPort string
    PostServiceHost string
    PostServicePort string
}

func GetConfig() (*Config,error){
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables",err.Error())
    }

	config := &Config{
        RateLimit:   getEnvInt("RATE_LIMIT",2),
        BucketSize: getEnvInt("BUCKET_SIZE",10),
        ServerPort: getEnv("PORT", ""),
        PostServiceHost: getEnv("POST_SERVICE_HOST_LOCAL",""),
        PostServicePort: getEnv("POST_SERVICE_PORT",""),
    }
	return config,nil
}

func getEnvInt(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists {
        result,err:= strconv.Atoi(value)
		if err == nil {
			return result
		}
    }
    return defaultValue
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

