package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Service struct {
	yt *YT
}

func (s *Service) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// todo: there is more channel infor https://www.googleapis.com/youtube/v3/videos?part=id%2C+snippet&id=4Y4YSpF6d6w&key=?
func (s *Service) Channel(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"channel": "todo",
	})
}

// const channelURL = "https://www.youtube.com/@i-hate-the-concert" // todo: get from request
func (s *Service) Videos(c *gin.Context) {
	channelURL := c.Param("channelURL")

	channelID, err := s.yt.GetChannelID(channelURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting channel id",
		})
		log.Fatal().Err(err).Msg("error getting channel id")
		return
	}

	videos, err := s.yt.GetVideos(channelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting videos",
		})
		log.Fatal().Err(err).Msg("error getting videos")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
	})
}
