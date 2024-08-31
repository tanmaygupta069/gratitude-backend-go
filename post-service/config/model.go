package config

type Config struct{
	Port string
	DynamoDBConfig DynamoDBConfig
}

type DynamoDBConfig struct{
	Port string
    Region string
	Host string
}



