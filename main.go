package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var (
	client    *redis.Client
	redisAddr string
	wsPort    int
)

func initRedisClient() error {
	// redis-12644.c1.us-east1-2.gce.cloud.redislabs.com:12644
	// localhost:6379
	// redis:6379
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func startServer() error {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", wsPort), nil)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.StringVar(&redisAddr, "redis", "", "Redis server address:port")
	flag.IntVar(&wsPort, "port", 8080, "Web server port")
	flag.Parse()

	if redisAddr == "" {
		log.Fatal("No redis client provided")
	}

	log.Printf("Connecting to %s\n", redisAddr)
	err := initRedisClient()
	if err != nil {
		log.Fatalf("Failed to connec to Redis client: %s\n", err.Error())
	}
	log.Println("Connected to Redis")

	log.Printf("Starting web server on %d\n", wsPort)
	err = startServer()
	if err != nil {
		log.Fatalf("Failed to start the server: %s\n", err.Error())
	}
}
