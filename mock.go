package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
)

func (s *Service) MockedVideos(c *gin.Context) {
	log.Trace().Msg("get videos")

	c.JSON(http.StatusOK, MockedData())
	return
}

func MockedData() gin.H {
	names := []string{"Я не піду на концерт", "Платівка", "Рокіт", "Поліфонік", "Звук", "Як написати пісню?", "Мій приватний канал", "Гітарна магія", "Музичний світ", "Український Рок", "Мій Метал", "Ааауууу"}
	videoTitles := []string{"Про акустичні хіти THE CURE", "Чому Кріст Новоселіч не грав у Foo Fighters?", "Продав душу дияволу! Перший в КЛУБІ 27. Роберт Джонсон.", "Що нам кричить Дейв Ґрол?", "Історія Американського Ідіота", "Where is my mind? Де мій розум? Де мої думки?", "Ліам побив Ноеля гітарою! Чому розпався OASIS?", "2 значення AROUND THE WORLD від Daft Punk!", "Фестиваль, який відбувся всупереч червоній владі", "Як Алберн ДВІЧІ змінив музику? Blur. Gorillaz."}
	return gin.H{
		"channelName":     names[randRange(0, len(names)-1)],
		"subscriberCount": randRange(1000, 10000),
		"viewCount":       randRange(100000, 10000000),
		"videoCount":      randRange(10, 100),
		"videos": []map[string]interface{}{
			{
				"Title":        videoTitles[randRange(0, len(videoTitles)-1)],
				"VideoID":      "NK31nH17lw8",
				"PublishedAt":  fmt.Sprintf("2024-08-0%dT06:00:11Z", 1),
				"Views":        randRange(1000, 30000),
				"Likes":        randRange(100, 3000),
				"CommentCount": randRange(10, 300),
			},
			{
				"Title":        videoTitles[randRange(0, len(videoTitles)-1)],
				"VideoID":      "A87k1FAMGQw",
				"PublishedAt":  fmt.Sprintf("2024-09-0%dT06:00:11Z", randRange(1, 9)),
				"Views":        randRange(1000, 30000),
				"Likes":        randRange(100, 3000),
				"CommentCount": randRange(10, 300),
			},
			{
				"Title":        videoTitles[randRange(0, len(videoTitles)-1)],
				"VideoID":      "sxWfd9lJAB0",
				"PublishedAt":  fmt.Sprintf("2024-10-0%dT06:00:11Z", randRange(1, 9)),
				"Views":        randRange(1000, 30000),
				"Likes":        randRange(100, 3000),
				"CommentCount": randRange(10, 300),
			},
			{
				"Title":        videoTitles[randRange(0, len(videoTitles)-1)],
				"VideoID":      "sxWfd9lJAB0",
				"PublishedAt":  fmt.Sprintf("2024-11-0%dT06:00:11Z", randRange(1, 9)),
				"Views":        randRange(1000, 30000),
				"Likes":        randRange(100, 3000),
				"CommentCount": randRange(10, 300),
			},
			{
				"Title":        videoTitles[randRange(0, len(videoTitles)-1)],
				"VideoID":      "sxWfd9lJAB0",
				"PublishedAt":  fmt.Sprintf("2024-12-%dT06:00:11Z", 31),
				"Views":        randRange(1000, 30000),
				"Likes":        randRange(100, 3000),
				"CommentCount": randRange(10, 300),
			},
		},
	}
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
