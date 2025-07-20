
### Go Chat App

A real-time WebSocket-based chat application built using Golang and the Gorilla WebSocket package.  
This application supports multi-client messaging and uses concurrency-safe `Hub` and `Client` structures with a minimal HTML/JavaScript frontend.

---

## Features

- Real-time 1:1 and group messaging
- Built using WebSockets (`gorilla/websocket`)
- Multi-client support via Go channels and goroutines
- Simple, responsive frontend (HTML + JavaScript)
- Message broadcast to all connected clients

---

## Tech Stack

| Layer     | Technology                  |
|-----------|-----------------------------|
| Backend   | Go 1.24, Gorilla WebSocket  |
| Frontend  | HTML, CSS, JavaScript       |
| Protocol  | WebSocket                   |
| Tools     | Go Modules, Git             |

---

## Project Structure

```

go-chat-app/
├── main.go         # HTTP server and WebSocket endpoints
├── hub.go          # Central hub to manage all clients
├── client.go       # Client connection handler
├── index.html      # Frontend UI for chat
├── go.mod          # Go module definition
└── .gitignore

````

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed (v1.18 or higher recommended)
- Git installed

### Run Locally

```bash
git clone https://github.com/YOUR_USERNAME/go-chat-app.git
cd go-chat-app

go mod tidy      # Install dependencies
go run .         # Run the server
````

Open your browser at: [http://localhost:8080](http://localhost:8080)

Open in multiple tabs to simulate multiple clients.

---

## Roadmap

* Add user names
* Support multiple chat rooms
* Store chat history (in-memory or database)
* Dockerize the application
* Deploy to a public server (Render, Railway, etc.)

---

