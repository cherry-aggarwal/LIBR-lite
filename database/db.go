package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cherry-aggarwal/libr/models"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func EnsureDatabaseExists(uri string) {
	fmt.Println("Controller init started")
	var dbName string = "libr"
	ctx := context.Background()
	var exists bool

	var newURI string = fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	Pool, err := pgxpool.New(ctx, newURI)
	if err != nil {
		fmt.Println("couldn't connect to postgres")
	}

	err = Pool.QueryRow(ctx, `
        SELECT EXISTS(
            SELECT 1
            FROM pg_catalog.pg_database
            WHERE datname = $1
        )`, "libr").Scan(&exists)
	if err != nil {
		log.Fatalf("checking of libr failed: %v", err)
	}
	if !exists {
		log.Printf("Database %q not found – creating...", dbName)
		if _, err := Pool.Exec(ctx, fmt.Sprintf(`CREATE DATABASE "%s"`, dbName)); err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		log.Printf("Database %q created.", dbName)
	} else {
		log.Printf("Database %q already exists.", dbName)
	}

	Pool, err = pgxpool.New(ctx, uri)
	if err != nil {
		log.Fatalf("Unable to connect to 'libr' database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS messages (
		id UUID PRIMARY KEY,
		content TEXT NOT NULL,
		timestamp BIGINT NOT NULL,
		status VARCHAR(10) NOT NULL
	)`
	_, err = Pool.Exec(ctx, createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}
}

func InitConnection() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}
	uri := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/libr?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
	)
	EnsureDatabaseExists(uri)

	var err error
	Pool, err = pgxpool.New(context.Background(), uri)
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
	log.Println("connected to db")

}

func InsertMessage(message models.Msg) (string, error) {
	query := "INSERT INTO messages(id,content,timestamp,status) VALUES ($1,$2,$3,$4)"
	_, err := Pool.Exec(context.Background(), query, message.MsgID, message.Content, message.TimeStamp, message.Status)
	if err != nil {
		fmt.Printf("error inserting message: %v", err)
		return "Error", err
	}
	return "Message Successfully Inserted", nil
}

func GetMessages(ts int64) []models.Msg {
	query := "SELECT * FROM messages WHERE timestamp = $1"
	rows, err := Pool.Query(context.Background(), query, ts)
	if err != nil {
		fmt.Printf("error getting messages from db: %v", err)
		return nil
	}
	defer rows.Close()

	var messages []models.Msg
	for rows.Next() {
		var message models.Msg
		if err := rows.Scan(&message.MsgID, &message.Content, &message.TimeStamp, &message.Status); err != nil {
			log.Fatalf("Error scanning row: %v", err)
			continue
		}
		if message.Status == "rejected" {
			continue
		}
		messages = append(messages, message)
	}
	fmt.Println(messages)
	return messages
}

// func main() {
// 	InitConnection()
// 	// var messgae1 = models.Msg{"hdsjsj", "fs", 34, "4vd"}
// 	// InsertMessage(messgae1)
// 	// GetMessages(34)
// 	defer Pool.Close()

// }
