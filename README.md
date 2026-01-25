# DevBook - Social Network in Go 🐹

![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)

**DevBook** is a full-featured social network application developed as the final project for the course *"Learn Golang from Scratch! Build a Complete Application"*.

This project was essential for consolidating advanced Go concepts, including concurrency, authentication (JWT), database integration, and the architectural separation between API and Frontend.

## 🚀 Features

The system is divided into two main modules: **API** (Backend) and **WebApp** (Frontend).

### 🔐 Authentication & Security
- User Registration and Login.
- Authentication via **JWT Tokens** (JSON Web Token).
- Custom Middlewares for route protection.

### 👥 Social Network
- **Posts:** Create, list, like, and unlike posts.
- **Interactions:** Follow and unfollow other users.
- **Profile:** Edit personal data and password management.

## 🛠️ Tech Stack

### Backend (API)
- **Language:** Go (Golang)
- **Routing:** Gorilla Mux
- **Database:** MySQL (running in a Docker container).
- **Authentication:** JWT-Go.

### Frontend (WebApp)
- **Rendering:** HTML/CSS/JavaScript.
- **API Consumption:** Direct Backend integration via HTTP Requests in Go.

### Prerequisites
- [Go](https://golang.org/dl/) installed.

```bash
cp .env.example .env
