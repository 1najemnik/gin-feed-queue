package services

import (
	"context"
	"log"

	"gin-feed-queue/models"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

var firestoreClient *firestore.Client

func SetFirestoreClient(client *firestore.Client) {
	firestoreClient = client
}

func GenerateUUID() string {
	return uuid.New().String()
}

func SaveNews(newsList []models.News) error {
	ctx := context.Background()

	for _, news := range newsList {
		exists, err := isNewsExistsByURL(news.URL)
		if err != nil {
			log.Printf("Failed to check existence for URL %s: %v", news.URL, err)
			continue
		}

		if exists {
			log.Printf("News with URL %s already exists. Skipping.", news.URL)
			continue
		}

		news.ID = GenerateUUID()
		_, err = firestoreClient.Collection("news").Doc(news.ID).Set(ctx, news)
		if err != nil {
			log.Printf("Failed to save news: %v", err)
			continue
		}

		log.Printf("News saved successfully: %s", news.URL)
	}

	return nil
}

func isNewsExistsByURL(url string) (bool, error) {
	ctx := context.Background()
	iter := firestoreClient.Collection("news").Where("url", "==", url).Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return doc.Exists(), nil
}

func GetAllNews(page, limit int) ([]models.News, error) {
	ctx := context.Background()
	var newsList []models.News

	offset := (page - 1) * limit
	log.Printf("Fetching news with Offset: %d, Limit: %d", offset, limit)

	iter := firestoreClient.Collection("news").
		OrderBy("created_at", firestore.Desc).
		Offset(offset).
		Limit(limit).
		Documents(ctx)

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			log.Printf("Error iterating Firestore documents: %v", err)
			return nil, err
		}

		var news models.News
		doc.DataTo(&news)
		log.Printf("Fetched news item: %+v", news)
		newsList = append(newsList, news)
	}

	log.Printf("Total news fetched: %d", len(newsList))
	return newsList, nil
}

func GetNewsByID(id string) (models.News, error) {
	ctx := context.Background()
	doc, err := firestoreClient.Collection("news").Doc(id).Get(ctx)
	if err != nil {
		return models.News{}, err
	}

	var news models.News
	doc.DataTo(&news)
	return news, nil
}

func DeleteNewsByID(id string) error {
	ctx := context.Background()
	_, err := firestoreClient.Collection("news").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Failed to delete document with ID %s: %v", id, err)
		return err
	}
	log.Printf("Successfully deleted document with ID %s", id)
	return nil
}

func UpdateNewsContent(id, content string) error {
	ctx := context.Background()
	_, err := firestoreClient.Collection("news").Doc(id).Update(ctx, []firestore.Update{
		{Path: "content", Value: content},
		{Path: "status", Value: "Processed"},
	})
	return err
}

func UpdateNewsStatus(id, status string) error {
	ctx := context.Background()
	_, err := firestoreClient.Collection("news").Doc(id).Update(ctx, []firestore.Update{
		{Path: "status", Value: status},
	})
	return err
}
