package main

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"
)

const LimitVideosAmount = 10

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

type Video struct {
	Title        string
	VideoID      string
	PublishedAt  string
	Views        uint64
	Likes        uint64
	Dislikes     uint64
	CommentCount uint64
}

func (y *YT) GetChannelID(channelURL string) (string, error) {
	chTitle, err := getChannelTitle(channelURL)
	if err != nil {
		return "", err
	}

	return y.getChannelIDByHandle(chTitle)
}

func getChannelTitle(channelURL string) (string, error) {
	if !strings.Contains(channelURL, "youtube.com") {
		// @i-hate-the-concert <- with @
		if channelURL[0] == '@' {
			return channelURL[1:], nil
		}
		// i-hate-the-concert <- without @
		return channelURL, nil
	}

	// if https://www.youtube.com/@i-hate-the-concert
	reHandle := regexp.MustCompile(`youtube\.com/@([a-zA-Z0-9_-]+)`)
	if match := reHandle.FindStringSubmatch(channelURL); match != nil {
		return match[1], nil
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
		return "", fmt.Errorf("error searching the channel: %v", err)
	}

	if len(response.Items) == 0 || response.Items[0].Id == nil || len(response.Items[0].Id.ChannelId) == 0 {
		return "", fmt.Errorf("the channel ID couldn't be found: %s", handle)
	}

	return response.Items[0].Id.ChannelId, nil
}

func (y *YT) GetVideos(channelID string) ([]*Video, error) {
	call := y.cli.Search.List([]string{"snippet"}).
		ChannelId(channelID).
		Type("video").
		Order("date"). // Sort by date (newest - first)
		MaxResults(LimitVideosAmount)

	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("error getting videos: %v", err)
	}

	vIDs := make([]string, 0, LimitVideosAmount)
	for _, item := range response.Items {
		vIDs = append(vIDs, item.Id.VideoId)
	}

	return y.getVideoStatistic(vIDs)
}

func (y *YT) getVideoStatistic(videoIDs []string) ([]*Video, error) {
	statsCall := y.cli.Videos.List([]string{"snippet", "statistics"}).
		Id(videoIDs...)

	statsResponse, err := statsCall.Do()
	if err != nil {
		return nil, fmt.Errorf("помилка отримання статистики відео: %v", err)
	}

	videos := make([]*Video, 0, LimitVideosAmount)
	for _, item := range statsResponse.Items {
		video := &Video{
			Title:        item.Snippet.Title,
			PublishedAt:  item.Snippet.PublishedAt,
			VideoID:      item.Id,
			Views:        item.Statistics.ViewCount,
			Likes:        item.Statistics.LikeCount,
			Dislikes:     item.Statistics.DislikeCount,
			CommentCount: item.Statistics.CommentCount,
		}
		videos = append(videos, video)
	}

	return videos, nil
}
