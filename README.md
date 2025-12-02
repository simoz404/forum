# Forum

A web-based communication platform built with **Go**, **SQLite**, and **Docker**. This project allows users to create posts, comment, categorize content, and interact via likes and dislikes, featuring a robust authentication system.

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)
![SQLite](https://img.shields.io/badge/SQLite-Database-003B57?style=flat&logo=sqlite)
![License](https://img.shields.io/badge/License-MIT-green)

## 📖 Table of Contents
- [About the Project](#about-the-project)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Database Structure](#database-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Running with Docker (Recommended)](#running-with-docker-recommended)
  - [Running Locally](#running-locally)
- [Usage](#usage)
- [Authors](#authors)

## About the Project
This project was created to understand the fundamentals of web development, including HTTP protocols, sessions, cookies, and database manipulation. It avoids modern frontend frameworks (React, Vue, etc.) to focus on raw HTML handling and backend logic.

## Features

### 🔐 Authentication & Security
* **Registration:** Users can sign up with an Email, Username, and Password.
* **Login/Logout:** Secure session management using **Cookies** (with expiration).
* **Encryption:** Passwords are hashed and encrypted using **bcrypt** before storage.
* **Validation:** Checks for unique emails and handles incorrect credentials gracefully.

### 💬 Communication
* **Posts:** Registered users can create posts with text content.
* **Comments:** Registered users can comment on posts.
* **Categories:** Posts can be associated with one or more categories for better organization.
* **Visibility:** Guests (non-registered users) can read posts and comments but cannot interact.

### 👍 Interactions
* **Likes/Dislikes:** Registered users can like or dislike posts and comments.
* **Counters:** Total likes and dislikes are visible to all users.

### 🔍 Filtering
Users can filter the feed by:
* Specific **Categories** (Subforums).
* **Created Posts:** View posts created by the logged-in user.
* **Liked Posts:** View posts liked by the logged-in user.

## Technology Stack
* **Backend:** Go (Golang)
* **Database:** SQLite3
* **Frontend:** HTML5, CSS3 (No Frameworks)
* **Containerization:** Docker
* **Key Packages:**
    * `github.com/mattn/go-sqlite3` (Database driver)
    * `golang.org/x/crypto/bcrypt` (Password hashing)
    * `github.com/google/uuid` (Session ID generation)

## Database Structure
The project uses a relational SQLite database. A typical Entity Relationship Diagram (ERD) for this project includes:
* **Users:** Stores ID, username, email, hashed password.
* **Posts:** Stores content, author ID, timestamps.
* **Comments:** Stores content, author ID, post ID.
* **Categories:** Stores category names.
* **Likes/Dislikes:** Join tables linking Users to Posts/Comments.

## Getting Started

### Prerequisites
* [Docker](https://docs.docker.com/get-docker/)
* [Go](https://go.dev/dl/) (Only if running locally without Docker)

### Installation
Clone the repository:
```bash
git clone [https://learn.zone01oujda.ma/git/saljaoui/forum](https://learn.zone01oujda.ma/git/saljaoui/forum)
cd forum
