package main

import (
	"context"
	"fmt"
	"regexp"

	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"
)

type YT struct {
	cli *yt.Service
}

func New(ctx context.Context, apiKey string) (*YT, error) {
	yts, err := yt.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &YT{cli: yts}, nil
}

func (y *YT) getChannelID(channelURL string) (string, error) {
	reHandle := regexp.MustCompile(`youtube\.com/@([a-zA-Z0-9_-]+)`)

	if match := reHandle.FindStringSubmatch(channelURL); match != nil {
		return y.getChannelIDByHandle(match[1])
	}

	return "", fmt.Errorf("unknown URL format: %s", channelURL)
}

func (y *YT) getChannelIDByHandle(handle string) (string, error) {
	call := y.cli.Search.List([]string{"snippet"}).
		Q(handle).
		Type("channel").
		MaxResults(1)

	response, err := call.Do()
	if err != nil {
		return "", fmt.Errorf("помилка пошуку каналу: %v", err)
	}

	// Перевіряємо, чи отримано результат і повертаємо channelId
	if len(response.Items) == 0 {
		return "", fmt.Errorf("канал не знайдено")
	}

	return response.Items[0].Id.ChannelId, nil
}
