package services

import (
	"errors"
	"gin-feed-queue/models"
	"time"

	"github.com/mmcdole/gofeed"
)

func FetchRSSFeeds(urls []string, days int) ([]models.News, error) {
	if len(urls) == 0 {
		return nil, errors.New("no RSS URLs provided")
	}

	parser := gofeed.NewParser()
	var newsList []models.News
	since := time.Now().AddDate(0, 0, -days)

	for _, url := range urls {
		feed, err := parser.ParseURL(url)
		if err != nil {
			continue
		}

		for _, item := range feed.Items {
			publishedTime, err := time.Parse(time.RFC1123Z, item.Published)
			if err != nil || publishedTime.Before(since) {
				continue
			}

			newsList = append(newsList, models.News{
				ID:        item.GUID,
				Title:     item.Title,
				Content:   item.Description,
				Status:    models.StatusAdded,
				CreatedAt: publishedTime,
				URL:       item.Link,
			})
		}
	}

	return newsList, nil
}
