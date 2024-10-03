package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err, "FDFGDFGD")
	}
	defer db.Close()

	// Initialize database tables
	initDB(db)
	http.HandleFunc("/api/users", usersHandler)
	// // Set up routes
	// http.HandleFunc("/", homehandler)
	// http.HandleFunc("/register", registerHandler)
	// http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/logout", logoutHandler)
	// http.HandleFunc("/posts", postsHandler)
	// http.HandleFunc("/posts/new", newPostHandler)
	// http.HandleFunc("/comments", commentsHandler)
	// http.HandleFunc("/like", likeHandler)
	fmt.Println("server runnig on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}
	tmp.Execute(w, nil)
}

func initDB(db *sql.DB) {
	// SQL statements to create tables
	tables := []string{`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`, `
        CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );`, `
        CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            post_id INTEGER,
            user_id INTEGER,
            content TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (post_id) REFERENCES posts(id),
            FOREIGN KEY (user_id) REFERENCES users(id)
        );`, `
        CREATE TABLE IF NOT EXISTS categories (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL UNIQUE
        );`, `
        CREATE TABLE IF NOT EXISTS post_categories (
            post_id INTEGER,
            category_id INTEGER,
            PRIMARY KEY (post_id, category_id),
            FOREIGN KEY (post_id) REFERENCES posts(id),
            FOREIGN KEY (category_id) REFERENCES categories(id)
        );`, `
        CREATE TABLE IF NOT EXISTS likes_dislikes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            post_id INTEGER,
            comment_id INTEGER,
            is_like BOOLEAN,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id),
            FOREIGN KEY (post_id) REFERENCES posts(id),
            FOREIGN KEY (comment_id) REFERENCES comments(id)
        );`}

	// Execute each CREATE TABLE statement
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Printf("Error creating table: %q\n", err)
			return
		}
	}

	log.Println("Database tables created successfully")
}

type user struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Use "-" to exclude from JSON output
	CreatedAt time.Time `json:"created_at"`
}

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	newUser := user{
// 		username: r.FormValue("username"),
// 		email:    r.FormValue("email"),
// 		password: r.FormValue("password"),
// 	}

// 	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", newUser.username, newUser.email, newUser.password)
// 	if err != nil {
// 		log.Printf("Error inserting user: %v", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := []user{}
    rows, err := db.Query("SELECT id, username, email, password, created_at FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var u user
        if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, u)
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsersHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Implement user login
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Implement user logout
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement post listing and filtering
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	// Implement new post creation
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement comment creation and listing
}

func likeHandler(w http.ResponseWriter, r *http.Request) {
	// Implement like/dislike functionality
}
