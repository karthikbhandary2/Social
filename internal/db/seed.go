package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/karthikbhandary2/Social/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy",
	"mallory", "olivia", "peggy", "trent", "victor", "walter", "xavier", "yvonne", "zane", "quincy",
	"sophia", "liam", "emma", "noah", "ava", "oliver", "isabella", "elijah", "mia", "james",
	"amelia", "lucas", "harper", "mason", "evelyn", "ethan", "abigail", "logan", "ella", "jacob",
	"scarlett", "alexander", "aiden", "chloe", "sebastian", "avery", "jack", "madison", "benjamin", "aria",
}

var titles = []string{
	"Go Basics Explained",
	"Mastering Slices in Go",
	"Concurrency Made Simple",
	"Building APIs with Go",
	"Understanding Goroutines",
	"Tips for Clean Code",
	"Structs and Interfaces",
	"Error Handling in Go",
	"Writing Tests in Go",
	"Go Modules Demystified",
	"Optimizing Go Performance",
	"Intro to Web Development",
	"Deploying Go Apps",
	"Managing Dependencies",
	"Design Patterns in Go",
	"Go for Beginners",
	"Logging Best Practices",
	"Go and Docker",
	"Handling JSON in Go",
	"Building CLI Tools",
}

var contents = []string{
	"Learn the basics of Go and start writing efficient programs today.",
	"Discover how slices work in Go and why they are so powerful.",
	"Understand the fundamentals of concurrency and how to use goroutines.",
	"Explore how to build RESTful APIs quickly and easily using Go.",
	"A step-by-step guide to mastering goroutines and channels.",
	"Simple tips to improve your code quality and readability in Go.",
	"Learn about structs and interfaces to create scalable applications.",
	"Best practices for handling errors effectively in Go applications.",
	"How to write and run unit tests to ensure your code is reliable.",
	"Demystify Go modules and manage dependencies with confidence.",
	"Optimize your Go programs for better performance and efficiency.",
	"A beginner’s guide to starting web development using Go frameworks.",
	"Learn how to deploy Go applications on cloud platforms with ease.",
	"Understand dependency management and keep your Go projects organized.",
	"Explore popular design patterns and their implementation in Go.",
	"Step into Go development with simple examples and clear explanations.",
	"Learn logging techniques to track and debug your Go applications.",
	"Use Docker to containerize and deploy your Go projects seamlessly.",
	"Handle JSON data efficiently for API development in Go.",
	"Create command-line tools in Go to automate your workflows.",
}

var tags = []string{
	"Go, Programming, Basics",
	"Go, Slices, Data Structures",
	"Concurrency, Goroutines, Programming",
	"APIs, Go, Web Development",
	"Goroutines, Concurrency, Channels",
	"Code Quality, Best Practices, Go",
	"Structs, Interfaces, Scalability",
	"Error Handling, Go, Best Practices",
	"Unit Testing, Go, Reliability",
	"Modules, Dependencies, Go",
	"Performance, Optimization, Go",
	"Web Development, Go, Beginners",
	"Deployment, Cloud, Go",
	"Dependency Management, Go, Best Practices",
	"Design Patterns, Go, Programming",
	"Go, Beginners, Programming",
	"Logging, Debugging, Go",
	"Docker, Containers, Go",
	"JSON, APIs, Go",
	"CLI, Automation, Go",
}

var comments = []string{
	"Great introduction! Very helpful.",
	"Awesome guide, I learned a lot about slices.",
	"Concurrency is tricky, but this made it simpler!",
	"Building APIs has never been easier. Thanks!",
	"Finally, I understand goroutines. Well done!",
	"Clean code is so important. Loved the tips.",
	"Structs and interfaces explained so clearly.",
	"Error handling is crucial—great advice here.",
	"Testing in Go just became my new favorite thing.",
	"Go modules are so confusing, but this helped.",
	"Performance tips saved me hours of debugging.",
	"Excited to start web development now. Thanks!",
	"Deploying Go apps is no longer intimidating.",
	"Dependency management tips were spot on.",
	"Design patterns made easy—excellent article!",
	"This is perfect for a beginner like me.",
	"Logging is so much easier now. Thanks!",
	"Docker and Go are a match made in heaven!",
	"JSON handling was always a pain—until now.",
	"My CLI tools are way better thanks to this.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)
	for _, user := range(users) {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}
	tx.Commit()
	posts := generatePosts(200, users)
	for _, post := range(posts) {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range(comments) {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error posting comment:", err)
			return
		}
	}
    log.Println("Seeding Complete!")
}

func generateUsers(num int) []*store.User {
    users := make([]*store.User, num)
    for i := 0; i < num; i++ {
        users[i] = &store.User{
            Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
            Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
        }
    }
    return users
}

func generatePosts(num int, users []*store.User) []*store.Post{
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID: user.ID,
			Title: titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
