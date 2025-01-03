package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"gin-feed-queue/models"
	"gin-feed-queue/services"

	"github.com/gin-gonic/gin"
)

func getAccessKey() string {
	return os.Getenv("ACCESS_KEY")
}

func RenderIndexPage(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 30
	news, err := services.GetAllNews(page, limit)
	if err != nil {
		log.Printf("Error fetching news: %v", err)
		c.HTML(http.StatusInternalServerError, "template.tmpl", gin.H{
			"Title":     "News List - Error",
			"Error":     "Failed to fetch news",
			"AccessKey": getAccessKey(),
		})
		return
	}

	log.Printf("Fetched %d news items for page %d", len(news), page)

	c.HTML(http.StatusOK, "template.tmpl", gin.H{
		"Title":           "Gin Feed Queue",
		"ContentTemplate": "index.tmpl",
		"News":            news,
		"CurrentPage":     page,
		"NextPage":        page + 1,
		"PreviousPage":    page - 1,
		"HasNextPage":     len(news) == limit,
		"HasPreviousPage": page > 1,
		"AccessKey":       getAccessKey(),
		"HasStatus":       HasStatus,
	})
}

func RenderEditNewsPage(c *gin.Context) {
	id := c.Param("id")
	news, err := services.GetNewsByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "template.tmpl", gin.H{
			"Title":     "Edit News - Error",
			"Error":     "Failed to fetch news",
			"AccessKey": getAccessKey(),
		})
		return
	}

	c.HTML(http.StatusOK, "template.tmpl", gin.H{
		"Title":           "Edit News",
		"ContentTemplate": "edit_news.tmpl",
		"News":            news,
		"AccessKey":       getAccessKey(),
	})
}

func EditNews(c *gin.Context) {
	id := c.Param("id")
	content := c.PostForm("content")

	err := services.UpdateNewsContent(id, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/?access_key="+getAccessKey())
}

func DeleteNews(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteNewsByID(id)
	if err != nil {
		log.Printf("Failed to delete news with ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/?access_key="+getAccessKey())
}

func PublishToTelegram(c *gin.Context) {
	id := c.Param("id")

	news, err := services.GetNewsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}

	if !services.HasStatus(news.Status, models.StatusProcessed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "News must be processed before publishing"})
		return
	}

	channelID := os.Getenv("TELEGRAM_CHANNEL_ID")
	err = services.PublishToTelegram(channelID, news.Content)
	if err != nil {
		log.Printf("Error publishing to Telegram: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to publish to Telegram",
			"details": err.Error(),
		})
		return
	}

	err = services.AddNewsStatus(id, models.StatusPublishedToTelegram)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news status"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/?access_key="+getAccessKey())
}

func HasStatus(status, bit int) bool {
	return (status & bit) != 0
}

func GetStatusStrings(status int) []string {
	var statuses []string

	if status&models.StatusProcessed != 0 {
		statuses = append(statuses, "Processed")
	}
	if status&models.StatusPublishedToTelegram != 0 {
		statuses = append(statuses, "Published to Telegram")
	}
	if status&models.StatusPublishedToFacebook != 0 {
		statuses = append(statuses, "Published to Facebook")
	}

	return statuses
}
