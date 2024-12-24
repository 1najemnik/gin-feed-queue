# Gin Feed Queue

Gin Feed Queue is a web application for managing news from RSS feeds, with features for editing and publishing to Telegram.

## Features
- Fetch news from specified RSS feeds.
- Display a list of news with bitwise statuses, allowing multiple statuses to be combined (e.g., `Processed` and `Published`).
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
    ACCESS_KEY=password
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
    curl -X POST "http://localhost:8080/api/rss?access_key=password&days=2"
    ```

## Usage
- **Home Page**: Displays a list of news items with actions.
- **Edit News**: Modify the content of news and update its status to `Processed`.
- **Publish to Telegram**: Available for news with the `Processed` status.

## Contributing

We welcome contributions! If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Open a pull request with a clear description of your changes.

Before submitting, ensure that your code adheres to the project's coding standards and passes all tests.

## License

This project is licensed under the **MIT License**.

## Author

This project was created and maintained by:

**[Ilya Gordon](https://github.com/1najemnik)**
Feel free to reach out at [ilyagdn@gmail.com](mailto:ilyagdn@gmail.com).
