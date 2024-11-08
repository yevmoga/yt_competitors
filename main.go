package main

import (
	"context"
	"fmt"

	"yt_competitors/configs"

	"github.com/rs/zerolog/log"
)

const channelURL = "https://www.youtube.com/@i-hate-the-concert" // todo: get from request

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

	channelId, err := yt.getChannelID(channelURL)
	if err != nil {
		log.Fatal().Err(err).Msg("error getting channel id")
	}

	fmt.Printf("Channel ID: %s\n", channelId)
}
