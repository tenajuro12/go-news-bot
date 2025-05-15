# Go News Bot

A Telegram bot that fetches news from RSS feeds, filters them by keywords, summarizes articles using OpenAI, and sends them to a Telegram channel.

## Features

- **RSS Feed Integration**: Automatically fetches news from multiple RSS feeds
- **Article Filtering**: Filters articles based on configurable keywords
- **AI-Powered Summaries**: Summarizes articles using OpenAI
- **Telegram Channel Publishing**: Posts news with summaries to your Telegram channel
- **Admin Commands**: Bot commands for administrators to manage the service

## Requirements

- Go 1.24+
- PostgreSQL
- Telegram Bot Token
- OpenAI API Key (optional for summarization)

## Installation

### Clone the repository

```bash
git clone https://github.com/yourusername/go-news-bot.git
cd go-news-bot
```

### Set up the database

Start PostgreSQL using Docker:

```bash
docker-compose -f docker-compose.dev.yml up -d
```

Run migrations (assuming you have `goose` installed):

```bash
goose -dir internal/storage/migrations postgres "user=postgres password=postgres dbname=news_db sslmode=disable" up
```

### Configuration

Create a `config.env` file with the following variables:

```
TELEGRAM_BOT_TOKEN=your_bot_token
TELEGRAM_CHANNEL_ID=your_channel_id
DATABASE_DSN=postgres://postgres:postgres@localhost:5432/news_db?sslmode=disable
FETCH_INTERVAL=30m
NOTIFICATION_INTERVAL=1h
FILTER_KEYWORDS=keyword1,keyword2
OPENAI_KEY=your_openai_key
OPENAI_PROMPT="Summarize this article in a concise way"
```

## Usage

### Build and run

```bash
go build -o news-bot cmd/main.go
./news-bot
```

### Bot Commands

- `/start` - Start the bot and get basic information

### Adding News Sources

Add news sources to your database:

```sql
INSERT INTO sources (name, feed_url) VALUES ('Source Name', 'https://example.com/rss');
```

## Architecture

The application consists of several components:

- **Fetcher**: Periodically fetches news from RSS feeds
- **Storage**: Stores articles and news sources in PostgreSQL
- **Notifier**: Sends articles to the Telegram channel
- **Summarizer**: Creates summaries of articles using OpenAI
- **Bot**: Handles Telegram bot commands

## Development

### Run in development mode

```bash
go run cmd/main.go
```

### Project Structure

- `cmd/` - Application entry points
- `internal/` - Internal application code
  - `bot/` - Telegram bot command handlers
  - `botkit/` - Bot framework utilities
  - `config/` - Application configuration
  - `fetcher/` - RSS feed fetching
  - `model/` - Data models
  - `notifier/` - Telegram channel notification
  - `source/` - RSS source implementation
  - `storage/` - Database operations
  - `summary/` - Article summarization

## License

This project is licensed under the MIT License - see the LICENSE file for details.
