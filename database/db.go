package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cherry-aggarwal/libr/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Pool *pgxpool.Pool

func EnsureDatabaseExists(uri string) {
	var dbName string = "libr"
	db, err := sql.Open("pgx", uri)
	if err != nil {
		log.Fatalf("failed opening default connection: %v", err)
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`, dbName,
	).Scan(&exists)
	if err != nil {
		log.Fatalf("failed checking database existence: %v", err)
	}
	if !exists {
		log.Printf("Database %q not found—creating...", dbName)
		_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbName))
		if err != nil {
			log.Fatalf("failed to create database: %v", err)
		}
		log.Printf("Database %q created.", dbName)
	} else {
		log.Printf("Database %q already exists.", dbName)
	}
}

func InitConnection() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}
	uri := os.Getenv("connectionString")

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
// 	GetMessages(34)
// 	defer pool.Close()

// }
