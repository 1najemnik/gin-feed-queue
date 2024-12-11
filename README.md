# Gin Feed Queue

Gin Feed Queue is a web application for managing news from RSS feeds, with features for editing and publishing to Telegram.

## Features
- Fetch news from specified RSS feeds.
- Display a list of news with statuses: `New`, `Processed`, `Published`.
- Edit news content.
- Publish news to Telegram.

## Requirements
- Go 1.19+
- Firebase with Firestore enabled
- Telegram bot with a token
- RSS feeds

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/1najemnik/gin-feed-queue
    cd gin-feed-queue
    ```
2. Configure the `.env` file:
    ```plaintext
    FIREBASE_DATABASE_URL=https://your-firebase-project.firebaseio.com
    TELEGRAM_BOT_TOKEN=your-telegram-bot-token
    TELEGRAM_CHANNEL_ID=@your_channel_id
    RSS_FEEDS=https://example.com/rss,https://another-rss.com/feed
    PORT=8080
    ```
3. Install dependencies:
    ```bash
    go mod tidy
    ```
4. Start the application:
    ```bash
    go run cmd/main.go
    ```

## API
- **POST `/api/rss`**: Fetch news from RSS feeds for the last few days.
    ```bash
    curl -X POST "http://localhost:8080/api/rss?days=2"
    ```

## Usage
- **Home Page**: Displays a list of news items with actions.
- **Edit News**: Modify the content of news and update its status to `Processed`.
- **Publish to Telegram**: Available for news with the `Processed` status.

## License
MIT
