
# Go Chat App

A real-time WebSocket-based chat application built using Golang and the Gorilla WebSocket package.  
This application supports multiple users and rooms with real-time message delivery via WebSockets and a simple HTML/JS frontend.

---

## Features

- Real-time messaging using WebSockets
- Multiple chat rooms with room-level isolation
- Display sender username with each message
- Multi-client support via Go channels and goroutines
- Lightweight, responsive frontend (HTML + JS)
- Clean separation of backend logic (`Hub`, `Client`, `main`)

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
├── hub.go          # Central hub to manage all chat rooms & clients
├── client.go       # WebSocket client connection logic
├── index.html      # Frontend chat interface
├── go.mod          # Go module definition
└── .gitignore

````

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed (v1.18 or higher)
- Git installed

### Run Locally

```bash
git clone https://github.com/YOUR_USERNAME/go-chat-app.git
cd go-chat-app

go mod tidy      # Install dependencies
go run .         # Run the server
````

Then open your browser at:
[http://localhost:8080](http://localhost:8080)

### Try It

1. Open in multiple tabs.
2. Use different usernames and enter the same or different room names.
3. Messages will now only appear to users in the same room.

---

## Roadmap

* [x] Add user names
* [x] Support multiple chat rooms
* [ ] Store chat history (in-memory or database)
* [ ] Dockerize the application
* [ ] Add "leave room" and "active users" features
* [ ] Deploy to Render / Railway / Vercel

---

