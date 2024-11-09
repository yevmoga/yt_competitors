package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"yt_competitors/configs"

	"github.com/rs/zerolog/log"
)

func main() {
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
	r.GET("/ping", s.Ping)
	r.GET("/channel", s.Channel)
	r.GET("/videos/:channelURL", s.Videos)

	log.Trace().Msg("ran router")
	err = r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal().Err(err).Msg("error run")
	}
}
