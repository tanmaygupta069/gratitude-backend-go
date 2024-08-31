package config

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"strconv"
)

func GetConfig() (*Config,error){
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

	config := &Config{
        Port: getEnv("PORT", ""),
		DynamoDBConfig: DynamoDBConfig{
			Port: getEnv("DYNAMO_PORT",""),
			Region: getEnv("DYNAMO_REGION","local"),
			Host: getEnv("DYNAMO_HOST",""),
		},
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

