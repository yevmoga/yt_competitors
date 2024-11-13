package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yt_competitors/configs"

	"github.com/rs/zerolog/log"
)

func main() {

	// Приклад назви та тегів двох відео
	titleA := []string{"go", "tutorial", "introduction"}
	titleB := []string{"golang", "guide", "beginner"}
	tagsA := []string{"go", "golang", "programming"}
	tagsB := []string{"golang", "code", "development"}

	// Обчислюємо схожість
	titleSim, tagsSim, err := calculateSimilarity(titleA, titleB, tagsA, tagsB)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Title Similarity: %.2f\n", titleSim)
	fmt.Printf("Tags Similarity: %.2f\n", tagsSim)

	log.Trace().Msg("starting yt_competitors")

	cfg, err := configs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("error in configs.New")
	}

	yt, err := New(context.Background(), cfg.ApiKey)
	if err != nil {
		log.Fatal().Err(err).Msg("error init youtube cli")
	}

	s := Service{yt: yt}

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/ping", s.Ping)
	r.GET("/channel", s.Channel)
	r.GET("/videos", s.MockedVideos)
	r.GET("/videos/:channelURL", s.Videos)

	log.Trace().Msg("ran router")
	err = r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal().Err(err).Msg("error run")
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
