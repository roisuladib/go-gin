package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/roisuladib/go-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server		*gin.Engine
	ctx			context.Context
	mongoClient	*mongo.Client
	redisClient	*redis.Client
)

func init()  {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Connect to RedisDB
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err = redisClient.Set(ctx, "test", "Welcome to golang with MongoDB and RedisDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client successfully connected")

	server = gin.Default()
}

func main()  {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't load config: %v", err));
	}

	defer mongoClient.Disconnect(ctx)

	value, err := redisClient.Get(ctx, "test").Result()

	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	router := server.Group("/api")
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	log.Fatal(server.Run(":" + config.Port))
}
