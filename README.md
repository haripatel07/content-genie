# Content-Genie

AI-powered content repurposing tool that transforms any article into multiple content formats: summaries, tweets, and LinkedIn posts.

## Features

- **Article Scraping**: Automatically extracts content from web articles
- **AI Content Generation**: Uses OpenAI GPT-3.5-turbo to generate:
  - Concise article summaries
  - Engaging Twitter threads (3 tweets)
  - Professional LinkedIn posts
- **Real-time Processing**: Background job processing with status updates
- **Modern Web UI**: Clean, responsive interface built with Next.js and DaisyUI
- **REST API**: Simple API for job submission and retrieval
- **Persistent Storage**: SQLite database for job history

## Tech Stack

### Backend
- **Go** - Main programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **SQLite** - Database
- **goquery** - HTML parsing and scraping
- **OpenAI Go SDK** - AI content generation

### Frontend
- **Next.js 15** - React framework
- **React 19** - UI library
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **DaisyUI** - Component library
- **react-hot-toast** - Notifications

## Prerequisites

- Go 1.24+ installed
- Node.js 18+ installed
- OpenAI API key

## Setup

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the backend directory:
   ```bash
   OPENAI_API_KEY=your_openai_api_key_here
   ```

4. The backend will run on `http://localhost:8080` by default.

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install Node.js dependencies:
   ```bash
   npm install
   ```

3. The frontend is configured to connect to `http://localhost:8080` (set in `.env.local`).

## Running the Application

### Development Mode

1. Start the backend:
   ```bash
   cd backend
   go run main.go
   ```

2. In a new terminal, start the frontend:
   ```bash
   cd frontend
   npm run dev
   ```

3. Open your browser to `http://localhost:3000`

### Production Build

1. Build the backend:
   ```bash
   cd backend
   go build -o content-genie-backend
   ```

2. Build the frontend:
   ```bash
   cd frontend
   npm run build
   npm start
   ```

## Usage

1. Open the web application in your browser
2. Enter an article URL in the input field
3. Click "Generate Content"
4. Wait for processing (status updates automatically)
5. View the generated summary, tweets, and LinkedIn post
6. Copy any content to clipboard using the copy buttons

## API Documentation

### POST /api/jobs
Create a new content generation job.

**Request Body:**
```json
{
  "url": "https://example.com/article"
}
```

**Response:**
```json
{
  "ID": 1,
  "original_url": "https://example.com/article",
  "status": "pending",
  "status_detail": "",
  "summary": "",
  "tweets": "",
  "linkedin_post": "",
  "CreatedAt": "2025-10-20T10:00:00Z"
}
```

### GET /api/jobs
Retrieve all jobs, ordered by creation date (newest first).

**Response:**
```json
[
  {
    "ID": 1,
    "original_url": "https://example.com/article",
    "status": "complete",
    "status_detail": "",
    "summary": "Article summary here...",
    "tweets": "[\"Tweet 1\", \"Tweet 2\", \"Tweet 3\"]",
    "linkedin_post": "LinkedIn post content here...",
    "CreatedAt": "2025-10-20T10:00:00Z"
  }
]
```

## Project Structure

```
content-genie/
├── backend/
│   ├── main.go              # Application entry point
│   ├── api/
│   │   ├── handlers.go      # HTTP request handlers
│   │   └── routes.go        # API route definitions
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── models/
│   │   └── job.go           # Data models
│   ├── services/
│   │   ├── openai.go        # OpenAI integration
│   │   ├── processor.go     # Job processing orchestration
│   │   └── scrapper.go      # Web scraping functionality
│   ├── go.mod
│   ├── go.sum
│   └── .env                 # Environment variables (create this)
├── frontend/
│   ├── app/
│   │   ├── layout.tsx       # Root layout
│   │   ├── page.tsx         # Main page component
│   │   └── globals.css      # Global styles
│   ├── components/          # React components (currently empty)
│   ├── package.json
│   ├── tailwind.config.ts
│   ├── next.config.ts
│   ├── tsconfig.json
│   └── .env.local           # Frontend environment variables
└── README.md
```

## Environment Variables

### Backend (.env)
- `OPENAI_API_KEY`: Your OpenAI API key (required)

### Frontend (.env.local)
- `NEXT_PUBLIC_API_URL`: Backend API URL (default: http://localhost:8080)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.