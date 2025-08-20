# Jarvis Chat WebApp

A React TypeScript chat interface for communicating with the Jarvis AI assistant.

## Features

- Modern, responsive chat interface
- Real-time messaging with the Jarvis backend
- Beautiful gradient design with chat bubbles
- Auto-scrolling to latest messages
- Loading indicators
- Error handling

## Development

To run the webapp:

```bash
# Install dependencies (if not already done)
yarn

# Start the development server
yarn dev
```

The app will be available at `http://localhost:5173`

## Backend Integration

This webapp connects to the Jarvis backend running on `http://localhost:7070`. 

To start the backend server:

```bash
# From the project root
cd ../../
go run ./cmd/server
```

Make sure you have the `GEMINI_API_KEY` environment variable set for the backend to work.

## API Endpoints

- `POST /chat` - Send a message to Jarvis and receive a response
- `GET /health` - Health check endpoint

## Technologies

- React 19
- TypeScript
- Vite
- Custom CSS for styling