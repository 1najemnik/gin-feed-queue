package controllers

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"gin-feed-queue/services"

	"github.com/gin-gonic/gin"
)

func FetchRSSNews(c *gin.Context) {
	daysParam := c.Query("days")
	days, err := strconv.Atoi(daysParam)
	if err != nil || days <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'days' parameter"})
		return
	}

	rssURLs := strings.Split(os.Getenv("RSS_FEEDS"), ",")
	if len(rssURLs) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No RSS feeds provided"})
		return
	}

	news, err := services.FetchRSSFeeds(rssURLs, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch RSS feeds"})
		return
	}

	err = services.SaveNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save news to Firebase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News fetched and saved", "count": len(news)})
}
