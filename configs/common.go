package configs

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	// ServerPort is the port the server listens on
	ServerPort int

	// PublicPath is the path to the public directory
	PublicPath string

	// RedisHost is the host of the redis server
	RedisHost     string
	RedisPort     string
	RedisDB       int
	RedisUsername string
	RedisPassword string

	// RedisKeyPrefix is the prefix for all keys stored in redis
	KeyPrefix = "apps:disruption"
)

// read config from .env file
func init() {
	// load .env file
	err := godotenv.Load(".env", ".env.local")
	if err != nil {
		fmt.Println("did not find .env file")
	}

	// load .env.local file
	err = godotenv.Overload(".env.local")
	if err != nil {
		fmt.Println("did not find .env.local file")
	}

	// SERVER: read expected values or fallback to defaults
	nextRawServerPort := os.Getenv("SERVER_PORT")
	nextServerPort, err := strconv.Atoi(nextRawServerPort)
	if err != nil {
		ServerPort = 3000
	} else {
		ServerPort = nextServerPort
	}

	nextPublicPath := os.Getenv("PUBLIC_PATH")
	if nextPublicPath != "" {
		PublicPath = nextPublicPath
	} else {
		PublicPath = "./public"
	}

	// REDIS: read expected values or fallback to defaults
	nextRedisHost := os.Getenv("REDIS_HOST")
	if nextRedisHost != "" {
		RedisHost = nextRedisHost
	} else {
		RedisHost = "localhost"
	}

	nextRedisPort := os.Getenv("REDIS_PORT")
	if nextRedisPort != "" {
		RedisPort = nextRedisPort
	} else {
		RedisPort = "6379"
	}

	nextRawRedisDB := os.Getenv("REDIS_DB")
	nextRedisDB, err := strconv.Atoi(nextRawRedisDB)
	if err != nil {
		RedisDB = 0
	} else {
		RedisDB = nextRedisDB
	}

	RedisUsername = os.Getenv("REDIS_USERNAME")
	RedisPassword = os.Getenv("REDIS_PASSWORD")

	fmt.Println("Config loaded:")
	fmt.Println("  ServerPort:", ServerPort)
	fmt.Println("  PublicPath:", PublicPath)
	fmt.Println("  RedisHost:", RedisHost)
	fmt.Println("  RedisPort:", RedisPort)
	fmt.Println("  RedisDB:", RedisDB)
	fmt.Println("  RedisUsername:", RedisUsername)
	fmt.Println("  RedisPassword:", strings.Repeat("*", len(RedisPassword)))
}
